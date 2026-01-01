
-- Create transactions table
CREATE TABLE transactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(255),
    role role NOT NULL,
    currency VARCHAR(10) NOT NULL DEFAULT 'ETB',
    inspection_period VARCHAR(255),
    item_category_id UUID NOT NULL,
    item_name VARCHAR(255) NOT NULL,
    item_description TEXT,
    price NUMERIC(12, 2) NOT NULL,
    shipping_method VARCHAR(255),
    seller_email VARCHAR(255) NOT NULL,
    seller_phone VARCHAR(50),
    buyer_email VARCHAR(255) NOT NULL,
    buyer_phone VARCHAR(50),
    status VARCHAR(255),
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now(),

    CONSTRAINT fk_item_category
        FOREIGN KEY (item_category_id)
        REFERENCES item_categories(id)
        ON DELETE CASCADE
);
