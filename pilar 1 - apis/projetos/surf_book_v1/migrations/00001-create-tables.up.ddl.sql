CREATE TABLE users (
    user_id UUID PRIMARY KEY,
    name        VARCHAR(300) NOT NULL,
    email       VARCHAR(300) NOT NULL,
    phone       VARCHAR(20),
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE tags (
    tag_id   UUID PRIMARY KEY,
    name     VARCHAR(900) NOT NULL,
    color    VARCHAR(6) NOT NULL,
    user_id  UUID NOT NULL REFERENCES users(user_id),
    deleted_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE notebooks (
    notebook_id UUID PRIMARY KEY,
    user_id     UUID REFERENCES users(user_id),
    icon        VARCHAR(50),
    name        VARCHAR(300) NOT NULL,
    image       VARCHAR(900),
    description VARCHAR(900),
    deleted_at  TIMESTAMP,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE meta_contents (
    content_id  UUID PRIMARY KEY,
    notebook_id UUID REFERENCES notebooks(notebook_id),
    user_id     UUID NOT NULL REFERENCES users(user_id),
    icon        VARCHAR(50),
    name        VARCHAR(300) NOT NULL,        
    deleted_at  TIMESTAMP,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE nodes_contents (
    node_id      UUID,
    content_id   UUID REFERENCES meta_contents(content_id),
    user_id      UUID NOT NULL REFERENCES users(user_id),
    notebook_id  UUID NOT NULL REFERENCES notebooks(notebook_id),
    deleted_at   TIMESTAMP,
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE meta_tag_contents (
    content_id   UUID REFERENCES meta_contents(content_id),
    tag_id       UUID NOT NULL REFERENCES tags(tag_id),
    notebook_id  UUID REFERENCES notebooks(notebook_id),
    user_id      UUID NOT NULL REFERENCES users(user_id),
    deleted_at   TIMESTAMP,
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (content_id, tag_id)
);
