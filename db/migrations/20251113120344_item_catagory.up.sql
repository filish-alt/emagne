
-- Create table for item categories (if not already created)
CREATE TABLE item_categories (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  name text UNIQUE NOT NULL,     -- car, clothes, electronics
  description text,
  created_at timestamptz DEFAULT now()
);