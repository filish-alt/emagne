CREATE TABLE category_attributes (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  category_id uuid NOT NULL REFERENCES item_categories(id) ON DELETE CASCADE,
  name text NOT NULL,                 -- vin, engine_number, size, color
  data_type varchar(20) NOT NULL,     -- string, number, boolean
  is_required boolean DEFAULT false,
  created_at timestamptz DEFAULT now(),
  UNIQUE(category_id, name)
);
