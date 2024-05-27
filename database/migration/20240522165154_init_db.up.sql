CREATE TABLE IF NOT EXISTS users (
    user_id uuid DEFAULT gen_random_uuid() NOT NULL CONSTRAINT users_pk PRIMARY KEY,
    username varchar(50) NOT NULL,
    email varchar(255) NOT NULL,
    password varchar(255) NOT NULL,
    role varchar(20) NOT NULL,
    created_at timestamp DEFAULT NOW(),
    updated_at timestamp DEFAULT NOW(),
    deleted_at timestamp DEFAULT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS user_username_idx ON users (username)
WHERE
    (deleted_at IS NULL);

CREATE UNIQUE INDEX IF NOT EXISTS user_email_idx ON users (email, role)
WHERE
    (deleted_at IS NULL);


CREATE INDEX IF NOT EXISTS user_role_idx ON users (role);
   
   
CREATE TABLE IF NOT EXISTS merchants (
    merchant_id uuid DEFAULT gen_random_uuid() NOT NULL CONSTRAINT merchants_pk PRIMARY KEY,
    name varchar(30) NOT NULL,
    merchant_category varchar(30) NOT NULL,
    image_url text NOT NULL,
    lat DECIMAL(9,6) NOT NULL,
    long DECIMAL(12,9) NOT NULL
);

CREATE INDEX IF NOT EXISTS merchant_name_idx ON merchants (name);
CREATE INDEX IF NOT EXISTS merchant_category_idx ON merchants (merchant_category);
CREATE INDEX IF NOT EXISTS merchant_lat_idx ON merchants (lat);
CREATE INDEX IF NOT EXISTS merchant_long_idx ON merchants (long);


CREATE TABLE IF NOT EXISTS products (
    product_id uuid DEFAULT gen_random_uuid() NOT NULL CONSTRAINT products_pk PRIMARY KEY,
    merchant_id uuid NOT NULL,
    name varchar(30) NOT NULL,
    product_category varchar(30) NOT NULL,
    price int NOT NULL,
    image_url text NOT NULL,
    CONSTRAINT product_merchant_fk foreign key (merchant_id) references merchants (merchant_id)
);

CREATE INDEX IF NOT EXISTS product_merchant_idx ON products (merchant_id);
CREATE INDEX IF NOT EXISTS product_name_idx ON products (name);
CREATE INDEX IF NOT EXISTS product_category_idx ON products (product_category);


CREATE TABLE IF NOT EXISTS invoices (
	id serial4 CONSTRAINT invoices_pk PRIMARY KEY,
    estimated_id uuid DEFAULT gen_random_uuid() NOT NULL,
    order_id uuid NULL,
    user_id uuid NOT NULL,
    CONSTRAINT invoice_user_fk foreign key (user_id) references users (user_id)
);

CREATE INDEX IF NOT EXISTS invoice_estimated_idx ON invoices (estimated_id);
CREATE INDEX IF NOT EXISTS invoice_order_idx ON invoices (order_id);
CREATE INDEX IF NOT EXISTS invoice_user_idx ON invoices (user_id);

CREATE TABLE IF NOT EXISTS invoice_marchants (
	id serial4 CONSTRAINT invoice_merchants_pk PRIMARY KEY,
	invoice_id int NOT NULL,
	merchant_id uuid NOT NULL,
	is_starting_point bool NOT NULL,
    CONSTRAINT invoice_marchant_invoice_fk foreign key (invoice_id) references invoices (id),
    CONSTRAINT invoice_marchant_merchant_fk foreign key (merchant_id) references merchants (merchant_id)
);

CREATE INDEX IF NOT EXISTS invoice_merchant_invoice_idx ON invoice_marchants (invoice_id);
CREATE INDEX IF NOT EXISTS invoice_merchant_merchant_idx ON invoice_marchants (merchant_id);
CREATE INDEX IF NOT EXISTS invoice_merchant_starting_idx ON invoice_marchants (is_starting_point);

CREATE TABLE IF NOT EXISTS invoice_merchant_products (
	id serial4 CONSTRAINT invoice_merchant_products_pk PRIMARY KEY,
	invoice_merchant_id int NOT NULL,
	product_id uuid NOT NULL,
	quantity int not null,
	price int not null,
	total_price int not null,
    CONSTRAINT invoice_marchant_product_im_fk foreign key (invoice_merchant_id) references invoice_marchants (id),
    CONSTRAINT invoice_marchant_product_p_fk foreign key (product_id) references products (product_id)
);

CREATE INDEX IF NOT EXISTS invoice_marchant_product_im_idx ON invoice_merchant_products (invoice_merchant_id);
CREATE INDEX IF NOT EXISTS invoice_marchant_product_p_idx ON invoice_merchant_products (product_id);