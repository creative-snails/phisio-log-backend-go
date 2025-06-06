-- =============================================
-- Clean Up (Drop statements)
-- =============================================
DROP TRIGGER IF EXISTS update_health_records_updated_at ON health_records;
DROP FUNCTION IF EXISTS update_updated_at_column();
DROP FUNCTION IF EXISTS validate_string_array();

-- Drop tables in reverse order of dependencies
DROP TABLE IF EXISTS medical_consultations;
DROP TABLE IF EXISTS symptoms;
DROP TABLE IF EXISTS health_records;
DROP TABLE IF EXISTS users;

DROP TYPE IF EXISTS progress_enum;
DROP TYPE IF EXISTS improvement_enum;
DROP TYPE IF EXISTS severity_enum;

-- =============================================
-- Extensions
-- =============================================
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- =============================================
-- Custom Types
-- =============================================
CREATE TYPE progress_enum AS ENUM ('open', 'closed', 'in-progress');
CREATE TYPE improvement_enum AS ENUM ('improving', 'stable', 'worsening', 'varying');
CREATE TYPE severity_enum AS ENUM ('mild', 'moderate', 'severe', 'variable');

-- =============================================
-- Early Functions
-- =============================================

CREATE OR REPLACE FUNCTION  validate_string_array(entries VARCHAR[])
RETURNS BOOLEAN AS $$
BEGIN
    IF entries IS NULL OR array_length(entries, 1) IS NULL THEN
        RETURN TRUE;
    END IF;

    RETURN (
        array_length(entries, 1) <= 50 AND 
        array_position(entries, '') IS NULL AND
        array_position(entries, NULL) IS NULL AND
        NOT EXISTS (
            SELECT 1
            FROM unnset(entries) AS e
            WHERE length(e) < 2 OR length(e) > 200
        )
    );
END;
$$ LANGUAGE plpgsql;

-- =============================================
-- Table Definitions
-- =============================================
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL CHECK (length(name) >= 2),
    email VARCHAR(255) NOT NULL UNIQUE CHECK (email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$')
);

CREATE TABLE health_records (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    -- user_id UUID NOT NULL,
    parent_record_id UUID NULL,
    description VARCHAR(2000) NOT NULL CHECK (length(description) >= 10), 
    progress progress_enum NOT NULL DEFAULT 'open',
    improvement improvement_enum NOT NULL DEFAULT 'stable',
    severity severity_enum NOT NULL DEFAULT 'variable',
    treatments_tried VARCHAR(200)[] CHECK (validate_string_array(treatments_tried)),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    -- CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_parent FOREIGN KEY (parent_record_id) REFERENCES health_records(id)
);

CREATE TABLE symptoms (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    health_record_id UUID NOT NULL,
    name VARCHAR(200) NOT NULL CHECK (length(name) >= 2),
    start_date TIMESTAMP,
    CONSTRAINT fk_health_record FOREIGN KEY (health_record_id)
        REFERENCES health_records(id)
);

CREATE TABLE medical_consultations (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    health_record_id UUID NOT NULL, 
    consultant VARCHAR(200) NOT NULL CHECK (length(consultant) >= 2),
    date TIMESTAMP NOT NULL,
    diagnosis VARCHAR(2000) NOT NULL CHECK (length(diagnosis) >= 10),
    follow_up_actions VARCHAR(200)[] CHECK (validate_string_array(follow_up_actions)),
    CONSTRAINT fk_health_record FOREIGN KEY (health_record_id) REFERENCES health_records(id)
);

-- =============================================
-- Triggers and Functions
-- =============================================
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE 'plpgsql';

CREATE TRIGGER update_health_records_updated_at
    BEFORE UPDATE ON health_records
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();
