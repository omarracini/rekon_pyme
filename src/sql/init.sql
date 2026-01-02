CREATE TABLE IF NOT EXISTS movements (
    id UUID PRIMARY KEY,
    account_id VARCHAR(50) NOT NULL,
    date TIMESTAMP NOT NULL,
    concept TEXT NOT NULL,
    amount BIGINT NOT NULL, -- Guardado en centavos
    currency VARCHAR(3) NOT NULL,
    type VARCHAR(10) NOT NULL, -- ABONO o CARGO
    is_conciliated BOOLEAN DEFAULT FALSE
);