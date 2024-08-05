

# reverse-uber-be

Backend project using Go.

## Checklist

- [ ] Set up flavors
- [x] Setting up database
- [ ] Add models
- [ ] Seed data
- [ ] Configurate pipeline
- [ ] Deploy


├── controllers         # contains api functions and main business logic
├── docs                # swagger files 
├── logs
├── middlewares         # request/response middlewares
│   └── validators      # data/request validators
├── models              
│   └── db              # collection models
├── routes              # router initialization
└── services            # general service & database actions