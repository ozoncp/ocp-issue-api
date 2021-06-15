-- +goose Up
-- +goose StatementBegin
CREATE TABLE issues
(
    id            bigserial PRIMARY KEY,
    class_room_id bigint    NOT NULL,
    task_id       bigint    NOT NULL,
    user_id       bigint    NOT NULL,
    deadline      timestamp NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE issues;
-- +goose StatementEnd
