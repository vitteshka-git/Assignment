# Task 15: Microservices Architecture Diagram & Technical Design

```
                        [ Client Request ]
                                │
                                ▼
                       ┌──────────────────┐
                       │   API Gateway    │ (Port 8080)
                       └────────┬─────────┘
                                │
         ┌──────────────────────┼──────────────────────┐
         │ (Sync REST)          │ (Sync REST)          │ (Sync REST)
         ▼                      ▼                      ▼
┌──────────────────┐  ┌──────────────────┐  ┌──────────────────┐
│ Employee Service │  │Department Service│  │   Auth Service   │
│   (Port 8081)    │  │   (Port 8082)    │  │   (Port 8083)    │
└────────┬─────────┘  └────────┬─────────┘  └────────┬─────────┘
         │                      │                      │
         ▼                      ▼                      ▼
┌──────────────────┐  ┌──────────────────┐  ┌──────────────────┐
│   employee_db    │  │  department_db   │  │     auth_db      │
└──────────────────┘  └──────────────────┘  └──────────────────┘
```
