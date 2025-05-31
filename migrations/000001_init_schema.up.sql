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

DROP TYPE IF EXISTS stage_enum;
DROP TYPE IF EXISTS severity_enum;
DROP TYPE IF EXISTS progression_enum;

-- =============================================
-- Extensions
-- =============================================
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- =============================================
-- Custom Types
-- =============================================
CREATE TYPE stage_enum AS ENUM ('open', 'closed', 'in-progress');
CREATE TYPE severity_enum AS ENUM ('mild', 'moderate', 'severe', 'variable');
CREATE TYPE progression_enum AS ENUM ('improving', 'stable', 'worsening', 'varying');
CREATE TYPE body_part_enum AS ENUM  (
        -- Front side body parts
        'head-front', 'neck-left-front', 'neck-right-front', 'shoulder-left-front',
        'shoulder-right-front', 'upper-arm-left-front', 'upper-arm-right-front',
        'elbow-left-front', 'elbow-right-front', 'forearm-left-front',
        'forearm-right-front', 'wrist-left-front', 'wrist-right-front',
        'hand-left-front', 'hand-right-front', 'chest-left', 'chest-right',
        'upper-abdomen-left', 'upper-abdomen-right', 'lower-abdomen-left',
        'lower-abdomen-right', 'hip-left-front', 'hip-right-front',
        'thigh-left-front', 'thigh-right-front', 'knee-left-front',
        'knee-right-front', 'lower-leg-left-front', 'lower-leg-right-front',
        'ankle-left-front', 'ankle-right-front', 'foot-left-front', 'foot-right-front',
        
        -- Back side body parts
        'head-back', 'shoulder-right-back', 'upper-arm-right-back', 'elbow-right-back',
        'forearm-right-back', 'wrist-right-back', 'hand-right-back', 'knee-right-back',
        'lower-leg-right-back', 'ankle-right-back', 'foot-right-back', 'buttocks-right',
        'middle-back-right', 'thigh-right-back', 'shoulder-blade-right', 'neck-right-back',
        'upper-spine-right', 'lower-spine-right', 'shoulder-left-back', 'upper-arm-left-back',
        'elbow-left-back', 'forearm-left-back', 'wrist-left-back', 'hand-left-back',
        'knee-left-back', 'lower-leg-left-back', 'ankle-left-back', 'foot-left-back',
        'buttocks-left-back', 'middle-back-left', 'thigh-left-back', 'shoulder-blade-left',
        'neck-left-back', 'upper-spine-left', 'lower-spine-left'
    );

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
            FROM unnest(entries) AS e
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
    parent_record_id UUID NULL,
    description VARCHAR(2000) NOT NULL CHECK (length(description) >= 10), 
    stage stage_enum NOT NULL DEFAULT 'open',
    severity severity_enum NOT NULL DEFAULT 'variable',
    progression progression_enum NOT NULL DEFAULT 'stable',
    treatments_tried VARCHAR(200)[] CHECK (validate_string_array(treatments_tried)),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_parent FOREIGN KEY (parent_record_id) REFERENCES health_records(id)
);

CREATE TABLE symptoms (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    health_record_id UUID NOT NULL,
    name VARCHAR(200) NOT NULL CHECK (length(name) >= 2),
    start_date TIMESTAMP,
    CONSTRAINT fk_health_record FOREIGN KEY (health_record_id)
        REFERENCES health_records(id)
        ON DELETE CASCADE
);

CREATE TABLE affected_parts (
    symptom_id UUID NOT NULL,
    body_part body_part_enum NOT NULL,
    state SMALLINT NOT NULL DEFAULT 1 CHECK (state BETWEEN 1 AND 4),
    PRIMARY KEY (symptom_id, body_part),
    CONSTRAINT fk_symptom FOREIGN KEY (symptom_id) REFERENCES symptoms(id)
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
