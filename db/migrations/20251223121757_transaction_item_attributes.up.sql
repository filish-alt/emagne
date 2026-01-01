
CREATE TABLE transaction_item_attributes (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  transaction_id uuid NOT NULL REFERENCES transactions(id) ON DELETE CASCADE,
  attribute_id uuid NOT NULL REFERENCES category_attributes(id),
  value text NOT NULL,
  created_at timestamptz DEFAULT now(),
  UNIQUE(transaction_id, attribute_id)
);
