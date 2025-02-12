	CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

	CREATE TABLE IF NOT EXISTS users(
		id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		email TEXT NOT NULL UNIQUE,
		password_hash TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT now(),
		updated_at TIMESTAMP DEFAULT now(),
		is_superuser BOOLEAN
	);

	CREATE TABLE IF NOT EXISTS articles(
		id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		title TEXT UNIQUE,
		description TEXT, 
		content TEXT,
		created_at TIMESTAMP DEFAULT now(),
		updated_at TIMESTAMP DEFAULT now(),
		public BOOLEAN
	);
    