
-- Create table echo
-- Migration ID: V1
-- Usage: Flyway
CREATE TABLE public.echo (
    id uuid NOT NULL,
    timestamp timestamp NOT NULL,
    value varchar(255) NOT NULL
)