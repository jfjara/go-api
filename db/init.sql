-- Crear tabla de usuarios con atributos JSONB
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    attributes JSONB,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Insertar usuario de ejemplo (solo si no existe)
INSERT INTO users (username, password, attributes)
VALUES (
    'jfjara',
    '$2a$10$8yo/toYc8u6jjhFmz1EbcOG9uAF6L7hnioVzlhGiouOwMOwAZ8/l2',
    '{"name": "JuanFran", "surname1": "Jara", "surname2": "Lopez"}'
)
ON CONFLICT (username) DO NOTHING;
