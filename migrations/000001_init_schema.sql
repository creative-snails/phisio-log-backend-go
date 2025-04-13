-- Drop tables in reverse order of dependencies
DROP TABLE IF EXISTS medical_consultations;
DROP TABLE IF EXISTS symptoms;
DROP TABLE IF EXISTS health_records;
DROP TABLE IF EXISTS users;

-- Drop types
DROP TYPE IF EXISTS progress_enum;
DROP TYPE IF EXISTS improvement_enum;
DROP TYPE IF EXISTS severity_enum;

-- Create extentions
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- ENUM types for the custom types
CREATE TYPE progress_enum AS ENUM ('open', 'closed', 'in-progress');
CREATE TYPE improvement_enum AS ENUM ('improving', 'stable', 'worsening', 'varying');
CREATE TYPE severity_enum AS ENUM ('mild', 'moderate', 'severe', 'variable');

-- Users table
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL CHECK (length(name) >= 2),
    email VARCHAR(255) NOT NULL CHECK (email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$')
);

-- Health Records table
CREATE TABLE health_records (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    parent_record_id UUID NULL,
    description VARCHAR(2000) NOT NULL CHECK (length(description) >= 10), 
    progress progress_enum NOT NULL DEFAULT 'open',
    improvement improvement_enum NOT NULL DEFAULT 'stable',
    severity severity_enum NOT NULL DEFAULT 'variable',
    treatments_tried VARCHAR(200)[] CHECK (
        array_length(treatments_tried, 1) IS NULL OR
        array_length(treatments_tried, 1) = 0 OR
        (SELECT bool_and(length(unnest) BETWEEN 2 AND 200) FROM unnest(treatments_tried))
    ),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_parent FOREIGN KEY (parent_record_id) REFERENCES health_records(id)
);

-- Symptoms table 
CREATE TABLE symptoms (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    health_record_id UUID NOT NULL,
    name VARCHAR(200) NOT NULL CHECK (length(name) >= 2),
    start_date TIMESTAMP,
    CONSTRAINT fk_health_record FOREIGN KEY (health_record_id)
        REFERENCES health_records(id)
);

-- Medical Consultations table
CREATE TABLE medical_consultations (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    health_record_id UUID NOT NULL, 
    consultant VARCHAR(200) NOT NULL CHECK (length(consultant) >= 2),
    date TIMESTAMP NOT NULL,
    diagnosis VARCHAR(2000) NOT NULL CHECK (length(diagnosis) >= 10),
    follow_up_actions VARCHAR(200)[] CHECK (
        array_length(follow_up_actions, 1) IS NULL OR
        array_length(follow_up_actions, 1) = 0 OR
        (SELECT bool_and(length(unnest) BETWEEN 2 AND 200) FROM unnest(follow_up_actions))
    ),
    CONSTRAINT fk_health_record FOREIGN KEY (health_record_id) REFERENCES health_records(id)
);

