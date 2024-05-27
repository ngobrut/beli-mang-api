DROP INDEX IF EXISTS user_username_idx;
DROP INDEX IF EXISTS user_email_idx;
DROP INDEX IF EXISTS user_role_idx;

DROP INDEX IF EXISTS merchant_name_idx;
DROP INDEX IF EXISTS merchant_category_idx;
DROP INDEX IF EXISTS merchant_lat_idx;
DROP INDEX IF EXISTS merchant_long_idx;

DROP INDEX IF EXISTS product_merchant_idx;
DROP INDEX IF EXISTS product_name_idx;
DROP INDEX IF EXISTS product_category_idx;

DROP INDEX IF EXISTS invoice_estimated_idx;
DROP INDEX IF EXISTS invoice_order_idx;
DROP INDEX IF EXISTS invoice_user_idx;

DROP INDEX IF EXISTS invoice_merchant_invoice_idx;
DROP INDEX IF EXISTS invoice_merchant_merchant_idx;
DROP INDEX IF EXISTS invoice_merchant_starting_idx;

DROP INDEX IF EXISTS invoice_marchant_product_im_idx;
DROP INDEX IF EXISTS invoice_marchant_product_p_idx;

DROP TABLE IF EXISTS invoice_merchant_products;
DROP TABLE IF EXISTS invoice_marchants;
DROP TABLE IF EXISTS invoices;
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS merchants;
DROP TABLE IF EXISTS users;