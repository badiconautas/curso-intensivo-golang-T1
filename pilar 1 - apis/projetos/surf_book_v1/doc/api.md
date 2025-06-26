# Surfbook V1

## Authententicacao
[BREVE]

## Entidades

## User (object)
+ user_id: `d2e6487c-b273-4c2e-9639-91d5a5d93d63` (string, required) - Identificador único (UUID)
+ name: Lucas Gomes (string, required)
+ email: lucas@email.com (string, required)
+ phone: +5511999999999 (string)
+ created_at: 2024-06-26T10:20:45Z (string, required) - Timestamp de criação
+ updated_at: 2024-06-26T10:20:45Z (string, required) - Timestamp de atualização

### Tag (object)
+ tag_id: `e58c12de-b9b4-4c12-9e3c-c5ecb6492dfe` (string, required) - Identificador único (UUID)
+ name: Importante (string, required)
+ color: FF0000 (string, required) - Cor em hexadecimal, ex: `FF0000`
+ user_id: `d2e6487c-b273-4c2e-9639-91d5a5d93d63` (string, required)
+ deleted_at: 2024-06-26T10:20:45Z (string, nullable)
+ created_at: 2024-06-26T10:20:45Z (string, required)
+ updated_at: 2024-06-26T10:20:45Z (string, required)

### Notebook (object)
+ notebook_id: `a9b1d668-776b-4980-8182-99eec97cfe1a` (string, required)
+ user_id: `d2e6487c-b273-4c2e-9639-91d5a5d93d63` (string, required)
+ icon: notebook (string, nullable)
+ name: Projetos 2025 (string, required)
+ image: https://exemplo.com/imagem.png (string, nullable)
+ description: Caderno para projetos de 2025 (string, nullable)
+ deleted_at: 2024-06-26T10:20:45Z (string, nullable)
+ created_at: 2024-06-26T10:20:45Z (string, required)
+ updated_at: 2024-06-26T10:20:45Z (string, required)

### MetaContent (object)
+ content_id: `f3f2635d-29f8-45ee-970a-540e3caab2e9` (string, required)
+ notebook_id: `a9b1d668-776b-4980-8182-99eec97cfe1a` (string, required)
+ user_id: `d2e6487c-b273-4c2e-9639-91d5a5d93d63` (string, required)
+ icon: file (string, nullable)
+ name: Anotações de Reunião (string, required)
+ deleted_at: 2024-06-26T10:20:45Z (string, nullable)
+ created_at: 2024-06-26T10:20:45Z (string, required)
+ updated_at: 2024-06-26T10:20:45Z (string, required)

### NodeContent (object)
+ node_id: `b3c847de-7a30-4c87-89c2-505fa84e733a` (string, required)
+ content_id: `f3f2635d-29f8-45ee-970a-540e3caab2e9` (string, required)
+ user_id: `d2e6487c-b273-4c2e-9639-91d5a5d93d63` (string, required)
+ notebook_id: `a9b1d668-776b-4980-8182-99eec97cfe1a` (string, required)
+ deleted_at: 2024-06-26T10:20:45Z (string, nullable)
+ created_at: 2024-06-26T10:20:45Z (string, required)
+ updated_at: 2024-06-26T10:20:45Z (string, required)

## Apis

Observações:

- Todos os endpoints requerem o header `user-id` com o UUID do usuário autenticado.
- Datas são sempre no formato ISO8601 UTC (ex: `"2024-06-26T10:20:45Z"`).
- Para erros, o retorno padrão é HTTP 400 e um JSON: `{ "error": "Mensagem" }`.
### Notebooks [Pronta]

#### List [GET /notebooks]

Listar Notebooks do Usuário

+ Headers
    + user-id: `d2e6487c-b273-4c2e-9639-91d5a5d93d63` (string, required) — UUID do usuário autenticado

+ Response 200 (application/json)

    + Body

            [
                {
                    "notebook_id": "a9b1d668-776b-4980-8182-99eec97cfe1a",
                    "user_id": "d2e6487c-b273-4c2e-9639-91d5a5d93d63",
                    "name": "Projetos 2025",
                    "description": "Caderno para projetos de 2025",
                    "icon": "notebook",
                    "image": "https://exemplo.com/imagem.png",
                    "deleted_at": null,
                    "created_at": "2024-06-26T10:20:45Z",
                    "updated_at": "2024-06-26T10:20:45Z"
                }
            ]

---

#### Find [GET /notebooks/{notebook_id}]

Buscar Notebook por ID
+ Parameters
    + notebook_id: `a9b1d668-776b-4980-8182-99eec97cfe1a` (string, required) — UUID do notebook

+ Headers
    + user-id: `d2e6487c-b273-4c2e-9639-91d5a5d93d63` (string, required)

+ Response 200 (application/json)

    + Body

            {
                "notebook_id": "a9b1d668-776b-4980-8182-99eec97cfe1a",
                "user_id": "d2e6487c-b273-4c2e-9639-91d5a5d93d63",
                "name": "Projetos 2025",
                "description": "Caderno para projetos de 2025",
                "icon": "notebook",
                "image": "https://exemplo.com/imagem.png",
                "deleted_at": null,
                "created_at": "2024-06-26T10:20:45Z",
                "updated_at": "2024-06-26T10:20:45Z"
            }

+ Response 400 (application/json)

    + Body

            {
                "error": "Notebook não encontrado: <mensagem>"
            }

---

#### Create [POST /notebooks]

Criar Notebook 
+ Headers
    + user-id: `d2e6487c-b273-4c2e-9639-91d5a5d93d63` (string, required) — UUID do usuário autenticado

+ Request (application/json)

    + Body

            {
                "name": "Projetos 2025",
                "description": "Caderno para projetos de 2025",
                "icon": "notebook",
                "image": "https://exemplo.com/imagem.png"
            }

+ Response 200 (application/json)

    + Body

            {
                "notebook_id": "a9b1d668-776b-4980-8182-99eec97cfe1a",
                "user_id": "d2e6487c-b273-4c2e-9639-91d5a5d93d63",
                "name": "Projetos 2025",
                "description": "Caderno para projetos de 2025",
                "icon": "notebook",
                "image": "https://exemplo.com/imagem.png",
                "deleted_at": null,
                "created_at": "2024-06-26T10:20:45Z",
                "updated_at": "2024-06-26T10:20:45Z"
            }

---


#### Update [PATCH /notebooks/{notebook_id}]
Atualizar Notebook 

+ Parameters
    + notebook_id: `a9b1d668-776b-4980-8182-99eec97cfe1a` (string, required)

+ Headers
    + user-id: `d2e6487c-b273-4c2e-9639-91d5a5d93d63` (string, required)

+ Request (application/json)

    + Body

            {
                "name": "Projetos 2026",
                "description": "Caderno atualizado",
                "icon": "notebook-edit",
                "image": "https://exemplo.com/nova-imagem.png"
            }

+ Response 200 (application/json)

    + Body

            {
                "notebook_id": "a9b1d668-776b-4980-8182-99eec97cfe1a",
                "user_id": "d2e6487c-b273-4c2e-9639-91d5a5d93d63",
                "name": "Projetos 2026",
                "description": "Caderno atualizado",
                "icon": "notebook-edit",
                "image": "https://exemplo.com/nova-imagem.png",
                "deleted_at": null,
                "created_at": "2024-06-26T10:20:45Z",
                "updated_at": "2024-06-27T10:20:45Z"
            }

---

#### Delete [DELETE /notebooks/{notebook_id}]
Deletar Notebook 

+ Parameters
    + notebook_id: `a9b1d668-776b-4980-8182-99eec97cfe1a` (string, required)

+ Headers
    + user-id: `d2e6487c-b273-4c2e-9639-91d5a5d93d63` (string, required)

+ Response 200 (application/json)

    + Body

            {
                "message": "Operação realizada com sucesso"
            }

---

### UserTags [ToDo]

#### List [GET /tags]

Listar Tags do Usuário

* Headers

  * user-id: `d2e6487c-b273-4c2e-9639-91d5a5d93d63` (string, required) — UUID do usuário autenticado

* Response 200 (application/json)

  * Body

    ```
      [
          {
              "tag_id": "e58c12de-b9b4-4c12-9e3c-c5ecb6492dfe",
              "name": "Importante",
              "color": "FF0000",
              "user_id": "d2e6487c-b273-4c2e-9639-91d5a5d93d63",
              "deleted_at": null,
              "created_at": "2024-06-26T10:20:45Z",
              "updated_at": "2024-06-26T10:20:45Z"
          }
      ]
    ```

#### Find [GET /tags/{tag_id}]

Buscar Tag por ID

* Parameters

  * tag_id: `e58c12de-b9b4-4c12-9e3c-c5ecb6492dfe` (string, required) — UUID da tag

* Headers

  * user-id: `d2e6487c-b273-4c2e-9639-91d5a5d93d63` (string, required)

* Response 200 (application/json)

  * Body

    ```
      {
          "tag_id": "e58c12de-b9b4-4c12-9e3c-c5ecb6492dfe",
          "name": "Importante",
          "color": "FF0000",
          "user_id": "d2e6487c-b273-4c2e-9639-91d5a5d93d63",
          "deleted_at": null,
          "created_at": "2024-06-26T10:20:45Z",
          "updated_at": "2024-06-26T10:20:45Z"
      }
    ```

* Response 400 (application/json)

  * Body

    ```
      {
          "error": "Tag não encontrada: <mensagem>"
      }
    ```

#### Create [POST /tags]

Criar Tag

* Headers

  * user-id: `d2e6487c-b273-4c2e-9639-91d5a5d93d63` (string, required)

* Request (application/json)

  * Body

    ```
      {
          "name": "Importante",
          "color": "FF0000"
      }
    ```

* Response 200 (application/json)

  * Body

    ```
      {
          "tag_id": "e58c12de-b9b4-4c12-9e3c-c5ecb6492dfe",
          "name": "Importante",
          "color": "FF0000",
          "user_id": "d2e6487c-b273-4c2e-9639-91d5a5d93d63",
          "deleted_at": null,
          "created_at": "2024-06-26T10:20:45Z",
          "updated_at": "2024-06-26T10:20:45Z"
      }
    ```

#### Update [PATCH /tags/{tag_id}]

Atualizar Tag

* Parameters

  * tag_id: `e58c12de-b9b4-4c12-9e3c-c5ecb6492dfe` (string, required)

* Headers

  * user-id: `d2e6487c-b273-4c2e-9639-91d5a5d93d63` (string, required)

* Request (application/json)

  * Body

    ```
      {
          "name": "Urgente",
          "color": "00FF00"
      }
    ```

* Response 200 (application/json)

  * Body

    ```
      {
          "tag_id": "e58c12de-b9b4-4c12-9e3c-c5ecb6492dfe",
          "name": "Urgente",
          "color": "00FF00",
          "user_id": "d2e6487c-b273-4c2e-9639-91d5a5d93d63",
          "deleted_at": null,
          "created_at": "2024-06-26T10:20:45Z",
          "updated_at": "2024-06-26T10:20:45Z"
      }
    ```

#### Delete [DELETE /tags/{tag_id}]

Deletar Tag

* Parameters

  * tag_id: `e58c12de-b9b4-4c12-9e3c-c5ecb6492dfe` (string, required)

* Headers

  * user-id: `d2e6487c-b273-4c2e-9639-91d5a5d93d63` (string, required)

* Response 200 (application/json)

  * Body

    ```
      {
          "message": "Operação realizada com sucesso"
      }
    ```

### MetaContents [ToDo]

#### List [GET /meta-contents]

Listar MetaContents do Usuário

* Headers

  * user-id: `d2e6487c-b273-4c2e-9639-91d5a5d93d63` (string, required)

* Response 200 (application/json)

  * Body

    ```
      [
          {
              "content_id": "f3f2635d-29f8-45ee-970a-540e3caab2e9",
              "notebook_id": "a9b1d668-776b-4980-8182-99eec97cfe1a",
              "user_id": "d2e6487c-b273-4c2e-9639-91d5a5d93d63",
              "icon": "file",
              "name": "Anotações de Reunião",
              "deleted_at": null,
              "created_at": "2024-06-26T10:20:45Z",
              "updated_at": "2024-06-26T10:20:45Z"
          }
      ]
    ```

#### Find [GET /meta-contents/{content_id}]

Buscar MetaContent por ID

* Parameters

  * content_id: `f3f2635d-29f8-45ee-970a-540e3caab2e9` (string, required)

* Headers

  * user-id: `d2e6487c-b273-4c2e-9639-91d5a5d93d63` (string, required)

* Response 200 (application/json)

  * Body

    ```
      {
          "content_id": "f3f2635d-29f8-45ee-970a-540e3caab2e9",
          "notebook_id": "a9b1d668-776b-4980-8182-99eec97cfe1a",
          "user_id": "d2e6487c-b273-4c2e-9639-91d5a5d93d63",
          "icon": "file",
          "name": "Anotações de Reunião",
          "deleted_at": null,
          "created_at": "2024-06-26T10:20:45Z",
          "updated_at": "2024-06-26T10:20:45Z"
      }
    ```

* Response 400 (application/json)

  * Body

    ```
      {
          "error": "MetaContent não encontrada: <mensagem>"
      }
    ```

#### Create [POST /meta-contents]

Criar MetaContent

* Headers

  * user-id: `d2e6487c-b273-4c2e-9639-91d5a5d93d63` (string, required)

* Request (application/json)

  * Body

    ```
      {
          "notebook_id": "a9b1d668-776b-4980-8182-99eec97cfe1a",
          "icon": "file",
          "name": "Anotações de Reunião"
      }
    ```

* Response 200 (application/json)

  * Body

    ```
      {
          "content_id": "f3f2635d-29f8-45ee-970a-540e3caab2e9",
          "notebook_id": "a9b1d668-776b-4980-8182-99eec97cfe1a",
          "user_id": "d2e6487c-b273-4c2e-9639-91d5a5d93d63",
          "icon": "file",
          "name": "Anotações de Reunião",
          "deleted_at": null,
          "created_at": "2024-06-26T10:20:45Z",
          "updated_at": "2024-06-26T10:20:45Z"
      }
    ```

#### Update [PATCH /meta-contents/{content_id}]

Atualizar MetaContent

* Parameters

  * content_id: `f3f2635d-29f8-45ee-970a-540e3caab2e9` (string, required)

* Headers

  * user-id: `d2e6487c-b273-4c2e-9639-91d5a5d93d63` (string, required)

* Request (application/json)

  * Body

    ```
      {
          "icon": "file-edit",
          "name": "Notas de Reunião"
      }
    ```

* Response 200 (application/json)

  * Body

    ```
      {
          "content_id": "f3f2635d-29f8-45ee-970a-540e3caab2e9",
          "notebook_id": "a9b1d668-776b-4980-8182-99eec97cfe1a",
          "user_id": "d2e6487c-b273-4c2e-9639-91d5a5d93d63",
          "icon": "file-edit",
          "name": "Notas de Reunião",
          "deleted_at": null,
          "created_at": "2024-06-26T10:20:45Z",
          "updated_at": "2024-07-01T12:00:00Z"
      }
    ```

#### Delete [DELETE /meta-contents/{content_id}]

Deletar MetaContent

* Parameters

  * content_id: `f3f2635d-29f8-45ee-970a-540e3caab2e9` (string, required)

* Headers

  * user-id: `d2e6487c-b273-4c2e-9639-91d5a5d93d63` (string, required)

* Response 200 (application/json)

  * Body

    ```
      {
          "message": "Operação realizada com sucesso"
      }
    ```

### NodeContents [ToDo]

#### List [GET /node-contents]

Listar NodeContents do Usuário

* Headers

  * user-id: `d2e6487c-b273-4c2e-9639-91d5a5d93d63` (string, required)

* Response 200 (application/json)

  * Body

    ```
      [
          {
              "node_id": "b3c847de-7a30-4c87-89c2-505fa84e733a",
              "content_id": "f3f2635d-29f8-45ee-970a-540e3caab2e9",
              "user_id": "d2e6487c-b273-4c2e-9639-91d5a5d93d63",
              "notebook_id": "a9b1d668-776b-4980-8182-99eec97cfe1a",
              "deleted_at": null,
              "created_at": "2024-06-26T10:20:45Z",
              "updated_at": "2024-06-26T10:20:45Z"
          }
      ]
    ```

#### Find [GET /node-contents/{node_id}]

Buscar NodeContent por ID

* Parameters

  * node_id: `b3c847de-7a30-4c87-89c2-505fa84e733a` (string, required)

* Headers

  * user-id: `d2e6487c-b273-4c2e-9639-91d5a5d93d63` (string, required)

* Response 200 (application/json)

  * Body

    ```
      {
          "node_id": "b3c847de-7a30-4c87-89c2-505fa84e733a",
          "content_id": "f3f2635d-29f8-45ee-970a-540e3caab2e9",
          "user_id": "d2e6487c-b273-4c2e-9639-91d5a5d93d63",
          "notebook_id": "a9b1d668-776b-4980-8182-99eec97cfe1a",
          "deleted_at": null,
          "created_at": "2024-06-26T10:20:45Z",
          "updated_at": "2024-06-26T10:20:45Z"
      }
    ```

* Response 400 (application/json)

  * Body

    ```
      {
          "error": "NodeContent não encontrada: <mensagem>"
      }
    ```

#### Create [POST /node-contents]

Criar NodeContent

* Headers

  * user-id: `d2e6487c-b273-4c2e-9639-91d5a5d93d63` (string, required)

* Request (application/json)

  * Body

    ```
      {
          "content_id": "f3f2635d-29f8-45ee-970a-540e3caab2e9",
          "notebook_id": "a9b1d668-776b-4980-8182-99eec97cfe1a"
      }
    ```

* Response 200 (application/json)

  * Body

    ```
      {
          "node_id": "b3c847de-7a30-4c87-89c2-505fa84e733a",
          "content_id": "f3f2635d-29f8-45ee-970a-540e3caab2e9",
          "user_id": "d2e6487c-b273-4c2e-9639-91d5a5d93d63",
          "notebook_id": "a9b1d668-776b-4980-8182-99eec97cfe1a",
          "deleted_at": null,
          "created_at": "2024-06-26T10:20:45Z",
          "updated_at": "2024-06-26T10:20:45Z"
      }
    ```

#### Update [PATCH /node-contents/{node_id}]

Atualizar NodeContent

* Parameters

  * node_id: `b3c847de-7a30-4c87-89c2-505fa84e733a` (string, required)

* Headers

  * user-id: `d2e6487c-b273-4c2e-9639-91d5a5d93d63` (string, required)

* Request (application/json)

  * Body

    ```
      {
          "content_id": "f3f2635d-29f8-45ee-970a-540e3caab2e9",
          "notebook_id": "a9b1d668-776b-4980-8182-99eec97cfe1a"
      }
    ```

* Response 200 (application/json)

  * Body

    ```
      {
          "node_id": "b3c847de-7a30-4c87-89c2-505fa84e733a",
          "content_id": "f3f2635d-29f8-45ee-970a-540e3caab2e9",
          "user_id": "d2e6487c-b273-4c2e-9639-91d5a5d93d63",
          "notebook_id": "a9b1d668-776b-4980-8182-99eec97cfe1a",
          "deleted_at": null,
          "created_at": "2024-06-26T10:20:45Z",
          "updated_at": "2024-06-26T10:20:45Z"
      }
    ```

#### Delete [DELETE /node-contents/{node_id}]

Deletar NodeContent

* Parameters

  * node_id: `b3c847de-7a30-4c87-89c2-505fa84e733a` (string, required)

* Headers

  * user-id: `d2e6487c-b273-4c2e-9639-91d5a5d93d63` (string, required)

* Response 200 (application/json)

  * Body

    ```
      {
          "message": "Operação realizada com sucesso"
      }
    ```

