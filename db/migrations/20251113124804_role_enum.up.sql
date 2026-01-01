
DROP TYPE IF EXISTS role CASCADE;
-- Create ENUM type for role
CREATE TYPE role AS ENUM (
	'Buyer',
	'Seller',
	'Broker');