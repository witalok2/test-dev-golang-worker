package repository

const (
	SqlCreateClient = `
		INSERT INTO client (
			name,
			last_name,
			contact,
			address,
			birthday,
			cpf
		)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	SqlUpdateClient = `
		UPDATE client
		SET
			name = $2,
			last_name = $3,
			contact = $4,
			address = $5,
			birthday = $6,
			cpf = $7,
			updated_at = NOW()
		WHERE
			id = $1
			AND deleted_at IS NULL
	`
)
