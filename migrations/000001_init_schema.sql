CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE progress_enum AS ENUM ('open', 'closed', 'in-progress');
CREATE TYPE improvement_enum AS ENUM ('improving', 'stable', 'worsening', 'varying');
CREATE TYPE severity_enum AS ENUM ('mild', 'moderate', 'severe', 'variable');

CREATE TABLE health_records (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    parent_record_id UUID NULL REFERENCES health_records(id),
    description TEXT NOT NULL, 
    progress progress_enum NOT NULL DEFAULT 'open',
    improvement improvement_enum NOT NULL DEFAULT 'stable',
    severity severity_enum NOT NULL DEFAULT 'variable',
    treatments_tried JSONB,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);