use store;

INSERT INTO stores (id,seller_user_id,name,created_at) VALUES
	 (1, 1,'edot store','2024-05-26 00:16:14'),
	 (2, 2,'another stores','2024-05-26 14:09:55');

INSERT INTO warehouses (id, store_id,name,status,created_at,updated_at) VALUES
	 (1, 1,'edot warehouse 1-1',1,'2024-05-26 00:17:40',NULL),
	 (2, 1,'edot warehouse 1-2',1,'2024-05-26 00:17:40',NULL),
	 (3, 2,'edot warehouse 2-1',1,'2024-05-26 00:17:40',NULL);

INSERT INTO warehouse_products (id, warehouse_id,product_id,stock) VALUES
	 (1, 1,'6651a98294344ee37278f6d6',100.00),
	 (2, 1,'6651a98294344ee37278f6d7',100.00),
	 (3, 1,'6651a98294344ee37278f6d8',100.00),
	 (4, 2,'6651a98294344ee37278f6d6',100.00),
	 (5, 2,'6651a98294344ee37278f6d7',100.00),
	 (6, 2,'6651a98294344ee37278f6d8',100.00);
INSERT INTO warehouse_products (id, warehouse_id,product_id,stock) VALUES
	 (7, 3,'6651a98294344ee37278f6d9',100.00),
	 (8, 3,'6651a98294344ee37278f6da',100.00);


