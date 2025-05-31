-- Drop triggers and functions
DROP TRIGGER IF EXISTS update_health_records_updated_at ON health_records;
DROP FUNCTION IF EXISTS update_updated_at_column();
DROP FUNCTION IF EXISTS validate_string_array();

-- Drop tables in reverse order of dependencies
DROP TABLE IF EXISTS medical_consultations;
DROP TABLE IF EXISTS affected_parts;
DROP TABLE IF EXISTS symptoms;
DROP TABLE IF EXISTS health_records;
-- DROP TABLE IF EXISTS users;

-- Drop custom types
DROP TYPE IF EXISTS stage_enum;
DROP TYPE IF EXISTS severity_enum;
DROP TYPE IF EXISTS progression_enum;
DROP TYPE IF EXISTS body_part_enum;