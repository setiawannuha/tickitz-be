DB_SOURCE=
MIGRATIONS_DIR=

# make migrate-init name="tbl_users"
# migrate create -dir ./migrations -ext sql user
migrate-init:
	migrate create -dir ${MIGRATIONS_DIR} -ext sql ${name}

# make migrate-up
# migrate -path ./migrations -database postgresql://postgres.vprrcdgjwrrbcbyjwbwu:x2DGOaAzPB6Hnsnz@aws-0-ap-southeast-1.pooler.supabase.com:6543/postgres -verbose up
migrate-up:
	migrate -path ${MIGRATIONS_DIR} -database ${DB_SOURCE} -verbose up

# make migrate-down
# migrate -path ./migrations -database postgresql://postgres.vprrcdgjwrrbcbyjwbwu:x2DGOaAzPB6Hnsnz@aws-0-ap-southeast-1.pooler.supabase.com:6543/postgres -verbose down
migrate-down:
	migrate -path ${MIGRATIONS_DIR} -database ${DB_SOURCE} -verbose down

# make migrate-fix
# migrate -path ./migrations -database postgresql://postgres.vprrcdgjwrrbcbyjwbwu:x2DGOaAzPB6Hnsnz@aws-0-ap-southeast-1.pooler.supabase.com:6543/postgres force 0
migrate-fix:
	migrate -path ${MIGRATIONS_DIR} -database ${DB_SOURCE} force 0
