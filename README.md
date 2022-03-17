
# ApiGomasterclass

The ApiGomasterclass api is used to retrieve main user data (first name, email, nickname) as well as data
concerning his video games, his progress, his library, his wishlist.


## API Reference

#### Get all games

```http
  GET /api/games
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `api_key` | `string` | **Required**. Your API key |

#### Get games

```http
  GET /api/games/${id}
```
| Parameter  | Type     | Description                          |
| :--------  | :------- | :--------------------------------    |
| `id`       | `string` | **Required**. Id of game to fetch    |

#### Delete games

```http
  GET /api/games/${id}
```
| Parameter  | Type     | Description                          |
| :--------  | :------- | :--------------------------------    |
| `id`       | `string` | **Required**. Id of game to Delete   |


#### Create games

```http
  POST /api/games
```

| Parameter    | Type     | Description                                 |
| :--------    | :------- | :--------------------------------           |
| `games_name` | `string` | **Required**. Name of item to create        |
| `mark`       | `string` | **Required**. mark of item to create        |
| `price`      | `string` | **Required**. Price of item to create       |
| `relase_date`| `string` | **Required**. Release date of item to create|


#### Update games

```http
  POST /api/games/${id}
```

| Parameter    | Type     | Description                                 |
| :--------    | :------- | :--------------------------------           |
| `id`         | `string` | **Required**. Id of game to update          |
| `games_name` | `string` | **Required**. Name of item to Update        |
| `mark`       | `string` | **Required**. mark of item to Update        |
| `price`      | `string` | **Required**. Price of item to create       |
| `relase_date`| `string` | **Required**. Release date of item to create|



## Installation

```
  git clone https://github.com/Aazarias/ApiGomasterclass.git
  go run main.go
```

## Authors

- [@Aazarias](https://github.com/Aazarias)
- Damasio Matthieu

