# Estimation Sheet - Test Results - REST API

### createManager
```json
HTTP/1.1 201 Created
Content-Type: application/json
Date: Tue, 24 Sep 2024 00:33:36 GMT
Content-Length: 200

{
  "user_id": "a72d61df-6b93-4557-bf30-bc9f08b823d6",
  "email": "john.doe@userland.com",
  "user_name": "john1234",
  "name": "John Doe",
  "user_type": "manager",
  "created_at": "2024-09-24T00:33:36Z",
  "updated_at": null
}
```

### createEstimator
```json
HTTP/1.1 201 Created
Content-Type: application/json
Date: Tue, 24 Sep 2024 00:34:12 GMT
Content-Length: 208

{
  "user_id": "adb7a7ee-9e5b-4d93-9f88-85a1abbea8db",
  "email": "marie.doe1110@userland.com",
  "user_name": "marie123",
  "name": "Marie Doe",
  "user_type": "estimator",
  "created_at": "2024-09-24T00:34:12Z",
  "updated_at": null
}
```

### listUsers
```json
HTTP/1.1 200 OK
Content-Type: application/json
Date: Tue, 24 Sep 2024 00:34:32 GMT
Content-Length: 420

{
  "users": [
    {
      "user_id": "a72d61df-6b93-4557-bf30-bc9f08b823d6",
      "email": "john.doe@userland.com",
      "user_name": "john1234",
      "name": "John Doe",
      "user_type": "manager",
      "created_at": "2024-09-24T00:33:36Z",
      "updated_at": null
    },
    {
      "user_id": "adb7a7ee-9e5b-4d93-9f88-85a1abbea8db",
      "email": "marie.doe1110@userland.com",
      "user_name": "marie123",
      "name": "Marie Doe",
      "user_type": "estimator",
      "created_at": "2024-09-24T00:34:12Z",
      "updated_at": null
    }
  ]
}
```
### createCompetence
```json
HTTP/1.1 201 Created
Content-Type: application/json
Date: Tue, 24 Sep 2024 00:35:06 GMT
Content-Length: 162

{
  "competence_id": "5506b538-0fd9-4d61-b58f-8ad68fd1070d",
  "code": "Tech Doc",
  "name": "Technical Documentation",
  "created_at": "2024-09-24T00:35:06Z",
  "updated_at": null
}
```
### updateTechDoc
```json
HTTP/1.1 200 OK
Content-Type: application/json
Date: Tue, 24 Sep 2024 00:36:24 GMT
Content-Length: 190

{
  "competence_id": "5506b538-0fd9-4d61-b58f-8ad68fd1070d",
  "code": "Tech Doc",
  "name": "Technical Documentation - Updated",
  "created_at": "2024-09-24T00:35:06Z",
  "updated_at": "2024-09-24T00:36:24Z"
}
```

### getCompetence
```json
HTTP/1.1 200 OK
Content-Type: application/json
Date: Tue, 24 Sep 2024 00:37:17 GMT
Content-Length: 190

{
  "competence_id": "5506b538-0fd9-4d61-b58f-8ad68fd1070d",
  "code": "Tech Doc",
  "name": "Technical Documentation - Updated",
  "created_at": "2024-09-24T00:35:06Z",
  "updated_at": "2024-09-24T00:36:24Z"
}
```

### listCompetences
```json
HTTP/1.1 200 OK
Content-Type: application/json
Date: Tue, 24 Sep 2024 00:38:01 GMT
Content-Length: 208

{
  "competences": [
    {
      "competence_id": "5506b538-0fd9-4d61-b58f-8ad68fd1070d",
      "code": "Tech Doc",
      "name": "Technical Documentation - Updated",
      "created_at": "2024-09-24T00:35:06Z",
      "updated_at": "2024-09-24T00:36:24Z"
    }
  ]
}
```

### createBaseline
```json
HTTP/1.1 201 Created
Content-Type: application/json
Date: Tue, 24 Sep 2024 00:38:41 GMT
Content-Length: 436

{
  "baseline_id": "1941a226-4f1f-450f-8831-ea3e74b480c1",
  "code": "RIT123456789",
  "review": 1,
  "title": "Logistics Cost \u0026 Time Management",
  "description": "This project will streamline our internal processes and increase overall efficiency",
  "duration": 12,
  "manager_id": "a72d61df-6b93-4557-bf30-bc9f08b823d6",
  "estimator_id": "adb7a7ee-9e5b-4d93-9f88-85a1abbea8db",
  "start_date": "2024-01-01",
  "created_at": "2024-09-24T00:38:41Z",
  "updated_at": null
}
```

### listBaselines
```json
HTTP/1.1 200 OK
Content-Type: application/json
Date: Tue, 24 Sep 2024 00:39:04 GMT
Content-Length: 497

{
  "baselines": [
    {
      "baseline_id": "1941a226-4f1f-450f-8831-ea3e74b480c1",
      "code": "RIT123456789",
      "review": 1,
      "title": "Logistics Cost \u0026 Time Management",
      "description": "This project will streamline our internal processes and increase overall efficiency",
      "duration": 12,
      "manager_id": "a72d61df-6b93-4557-bf30-bc9f08b823d6",
      "manager": "John Doe",
      "estimator_id": "adb7a7ee-9e5b-4d93-9f88-85a1abbea8db",
      "estimator": "Marie Doe",
      "start_date": "2024-01-01",
      "created_at": "2024-09-24T00:38:41Z",
      "updated_at": null
    }
  ]
}
```

### createBaselineWithError
```json
HTTP/1.1 422 Unprocessable Entity
Content-Type: application/json
Date: Tue, 24 Sep 2024 00:39:22 GMT
Content-Length: 161

{
  "status_code": 422,
  "error": "Unprocessable Entity",
  "message": {
    "invalid_payload": [
      {
        "field": "estimator_id",
        "error": "EstimatorID must be a valid version 4 UUID"
      }
    ]
  }
}
```

### updateBaseline
```json
HTTP/1.1 200 OK
Content-Type: application/json
Date: Tue, 24 Sep 2024 00:45:17 GMT
Content-Length: 454

{
  "baseline_id": "1941a226-4f1f-450f-8831-ea3e74b480c1",
  "code": "RIT123456789",
  "review": 1,
  "title": "Logistics Cost \u0026 Time Management",
  "description": "This project will streamline our internal processes and increase overall efficiency",
  "duration": 12,
  "manager_id": "a72d61df-6b93-4557-bf30-bc9f08b823d6",
  "estimator_id": "adb7a7ee-9e5b-4d93-9f88-85a1abbea8db",
  "start_date": "2024-01-01",
  "created_at": "2024-09-24T00:38:41Z",
  "updated_at": "2024-09-24T00:45:17Z"
}
```

### getBaseline
```json
HTTP/1.1 200 OK
Content-Type: application/json
Date: Tue, 24 Sep 2024 00:45:56 GMT
Content-Length: 499

{
  "baseline_id": "1941a226-4f1f-450f-8831-ea3e74b480c1",
  "code": "RIT123456789",
  "review": 1,
  "title": "Logistics Cost \u0026 Time Management",
  "description": "This project will streamline our internal processes and increase overall efficiency",
  "duration": 12,
  "manager_id": "a72d61df-6b93-4557-bf30-bc9f08b823d6",
  "manager": "John Doe",
  "estimator_id": "adb7a7ee-9e5b-4d93-9f88-85a1abbea8db",
  "estimator": "Marie Doe",
  "start_date": "2024-01-01",
  "created_at": "2024-09-24T00:38:41Z",
  "updated_at": "2024-09-24T00:45:17Z"
}
```

### createCostPO
```json
HTTP/1.1 201 Created
Content-Type: application/json
Date: Tue, 24 Sep 2024 00:46:43 GMT
Content-Length: 564

{
  "cost_id": "7b086193-87bc-49d6-8333-10b58735f1ad",
  "baseline_id": "1941a226-4f1f-450f-8831-ea3e74b480c1",
  "cost_type": "one_time",
  "description": "Mão de obra do PO",
  "comment": "estimativa do PO",
  "amount": 180000,
  "currency": "BRL",
  "tax": 0,
  "apply_inflation": true,
  "cost_allocations": [
    {
      "year": 2024,
      "month": 1,
      "amount": 30000
    },
    {
      "year": 2024,
      "month": 2,
      "amount": 30000
    },
    {
      "year": 2024,
      "month": 3,
      "amount": 30000
    },
    {
      "year": 2024,
      "month": 4,
      "amount": 30000
    },
    {
      "year": 2024,
      "month": 5,
      "amount": 30000
    },
    {
      "year": 2024,
      "month": 6,
      "amount": 30000
    }
  ],
  "created_at": "2024-09-24T00:46:43Z",
  "updated_at": null
}
```

### createCostConsulting
```json
HTTP/1.1 201 Created
Content-Type: application/json
Date: Tue, 24 Sep 2024 00:49:55 GMT
Content-Length: 429

{
  "cost_id": "75d97caf-a81c-4604-8d2e-b85d88b5b626",
  "baseline_id": "1941a226-4f1f-450f-8831-ea3e74b480c1",
  "cost_type": "one_time",
  "description": "External Consulting",
  "comment": "estimativa de consultoria externa",
  "amount": 80000,
  "currency": "EUR",
  "tax": 23.1,
  "apply_inflation": false,
  "cost_allocations": [
    {
      "year": 2024,
      "month": 4,
      "amount": 30000
    },
    {
      "year": 2024,
      "month": 6,
      "amount": 50000
    }
  ],
  "created_at": "2024-09-24T00:49:55Z",
  "updated_at": null
}
```

### getCostsByBaselineId
```json
HTTP/1.1 200 OK
Content-Type: application/json
Date: Tue, 24 Sep 2024 00:50:33 GMT
Content-Length: 1005

{
  "costs": [
    {
      "cost_id": "75d97caf-a81c-4604-8d2e-b85d88b5b626",
      "baseline_id": "1941a226-4f1f-450f-8831-ea3e74b480c1",
      "cost_type": "one_time",
      "description": "External Consulting",
      "comment": "estimativa de consultoria externa",
      "amount": 80000,
      "currency": "EUR",
      "tax": 23.1,
      "apply_inflation": false,
      "cost_allocations": [
        {
          "year": 2024,
          "month": 4,
          "amount": 30000
        },
        {
          "year": 2024,
          "month": 6,
          "amount": 50000
        }
      ],
      "created_at": "2024-09-24T00:49:55Z",
      "updated_at": null
    },
    {
      "cost_id": "7b086193-87bc-49d6-8333-10b58735f1ad",
      "baseline_id": "1941a226-4f1f-450f-8831-ea3e74b480c1",
      "cost_type": "one_time",
      "description": "Mão de obra do PO",
      "comment": "estimativa do PO",
      "amount": 180000,
      "currency": "BRL",
      "tax": 0,
      "apply_inflation": true,
      "cost_allocations": [
        {
          "year": 2024,
          "month": 1,
          "amount": 30000
        },
        {
          "year": 2024,
          "month": 2,
          "amount": 30000
        },
        {
          "year": 2024,
          "month": 3,
          "amount": 30000
        },
        {
          "year": 2024,
          "month": 4,
          "amount": 30000
        },
        {
          "year": 2024,
          "month": 5,
          "amount": 30000
        },
        {
          "year": 2024,
          "month": 6,
          "amount": 30000
        }
      ],
      "created_at": "2024-09-24T00:46:43Z",
      "updated_at": null
    }
  ]
}
```
### updateCostConsulting
```json
HTTP/1.1 200 OK
Content-Type: application/json
Date: Tue, 24 Sep 2024 00:52:48 GMT
Content-Length: 456

{
  "cost_id": "75d97caf-a81c-4604-8d2e-b85d88b5b626",
  "baseline_id": "1941a226-4f1f-450f-8831-ea3e74b480c1",
  "cost_type": "one_time",
  "description": "External Consulting",
  "comment": "estimativa de consultoria externa atualizada",
  "amount": 80000,
  "currency": "USD",
  "tax": 23,
  "apply_inflation": false,
  "cost_allocations": [
    {
      "year": 2024,
      "month": 7,
      "amount": 50000
    },
    {
      "year": 2024,
      "month": 8,
      "amount": 30000
    }
  ],
  "created_at": "2024-09-24T00:49:55Z",
  "updated_at": "2024-09-24T00:52:48Z"
}
```
### createEffort
```json
HTTP/1.1 201 Created
Content-Type: application/json
Date: Wed, 26 Feb 2025 02:09:59 GMT
Content-Length: 649
Connection: close

{
  "effort_id": "2b3c485a-eccf-41ab-8538-9aaca68d97a6",
  "baseline_id": "612ecf4d-9342-4fbd-a5fe-15f7a9814950",
  "competence_id": "bdb9f625-80de-4698-b51c-389bed7954f5",
  "competence_code": "Tech Doc",
  "competence_name": "Technical Documentation",
  "comment": "considerado a Simone na atividade",
  "hours": 160,
  "effort_allocations": [
    {
      "year": 2024,
      "month": 1,
      "hours": 20
    },
    {
      "year": 2024,
      "month": 2,
      "hours": 20
    },
    {
      "year": 2024,
      "month": 3,
      "hours": 20
    },
    {
      "year": 2024,
      "month": 4,
      "hours": 20
    },
    {
      "year": 2024,
      "month": 5,
      "hours": 20
    },
    {
      "year": 2024,
      "month": 6,
      "hours": 20
    },
    {
      "year": 2024,
      "month": 7,
      "hours": 20
    },
    {
      "year": 2024,
      "month": 8,
      "hours": 20
    }
  ],
  "created_at": "2025-02-26T02:09:59Z",
  "updated_at": null
}
```
### updateEffort
```json
HTTP/1.1 200 OK
Content-Type: application/json
Date: Wed, 26 Feb 2025 02:12:47 GMT
Content-Length: 674
Connection: close

{
  "effort_id": "2b3c485a-eccf-41ab-8538-9aaca68d97a6",
  "baseline_id": "612ecf4d-9342-4fbd-a5fe-15f7a9814950",
  "competence_id": "bdb9f625-80de-4698-b51c-389bed7954f5",
  "competence_code": "Tech Doc",
  "competence_name": "Technical Documentation",
  "comment": "considerado Simone e Regina na atividade",
  "hours": 160,
  "effort_allocations": [
    {
      "year": 2024,
      "month": 1,
      "hours": 20
    },
    {
      "year": 2024,
      "month": 2,
      "hours": 20
    },
    {
      "year": 2024,
      "month": 3,
      "hours": 20
    },
    {
      "year": 2024,
      "month": 4,
      "hours": 20
    },
    {
      "year": 2024,
      "month": 5,
      "hours": 20
    },
    {
      "year": 2024,
      "month": 6,
      "hours": 20
    },
    {
      "year": 2024,
      "month": 7,
      "hours": 20
    },
    {
      "year": 2024,
      "month": 8,
      "hours": 20
    }
  ],
  "created_at": "2025-02-26T02:09:59Z",
  "updated_at": "2025-02-26T02:12:47Z"
}
```
### getEffort
```json
HTTP/1.1 200 OK
Content-Type: application/json
Date: Wed, 26 Feb 2025 02:13:23 GMT
Content-Length: 674
Connection: close

{
  "effort_id": "2b3c485a-eccf-41ab-8538-9aaca68d97a6",
  "baseline_id": "612ecf4d-9342-4fbd-a5fe-15f7a9814950",
  "competence_id": "bdb9f625-80de-4698-b51c-389bed7954f5",
  "competence_code": "Tech Doc",
  "competence_name": "Technical Documentation",
  "comment": "considerado Simone e Regina na atividade",
  "hours": 160,
  "effort_allocations": [
    {
      "year": 2024,
      "month": 1,
      "hours": 20
    },
    {
      "year": 2024,
      "month": 2,
      "hours": 20
    },
    {
      "year": 2024,
      "month": 3,
      "hours": 20
    },
    {
      "year": 2024,
      "month": 4,
      "hours": 20
    },
    {
      "year": 2024,
      "month": 5,
      "hours": 20
    },
    {
      "year": 2024,
      "month": 6,
      "hours": 20
    },
    {
      "year": 2024,
      "month": 7,
      "hours": 20
    },
    {
      "year": 2024,
      "month": 8,
      "hours": 20
    }
  ],
  "created_at": "2025-02-26T02:09:59Z",
  "updated_at": "2025-02-26T02:12:47Z"
}
```
### listEfforts
```json
HTTP/1.1 200 OK
Content-Type: application/json
Date: Wed, 26 Feb 2025 02:16:03 GMT
Content-Length: 688
Connection: close

{
  "efforts": [
    {
      "effort_id": "2b3c485a-eccf-41ab-8538-9aaca68d97a6",
      "baseline_id": "612ecf4d-9342-4fbd-a5fe-15f7a9814950",
      "competence_id": "bdb9f625-80de-4698-b51c-389bed7954f5",
      "competence_code": "Tech Doc",
      "competence_name": "Technical Documentation",
      "comment": "considerado Simone e Regina na atividade",
      "hours": 160,
      "effort_allocations": [
        {
          "year": 2024,
          "month": 1,
          "hours": 20
        },
        {
          "year": 2024,
          "month": 2,
          "hours": 20
        },
        {
          "year": 2024,
          "month": 3,
          "hours": 20
        },
        {
          "year": 2024,
          "month": 4,
          "hours": 20
        },
        {
          "year": 2024,
          "month": 5,
          "hours": 20
        },
        {
          "year": 2024,
          "month": 6,
          "hours": 20
        },
        {
          "year": 2024,
          "month": 7,
          "hours": 20
        },
        {
          "year": 2024,
          "month": 8,
          "hours": 20
        }
      ],
      "created_at": "2025-02-26T02:09:59Z",
      "updated_at": "2025-02-26T02:12:47Z"
    }
  ]
}
```

### createPlanBP
```json
HTTP/1.1 201 Created
Content-Type: application/json
Date: Tue, 24 Sep 2024 00:54:59 GMT
Content-Length: 617

{
  "plan_id": "921350d4-b9ce-443e-b284-b0287213919b",
  "code": "BP 2025",
  "name": "Business Plan 2025",
  "assumptions": [
    {
      "year": 2024,
      "inflation": 4,
      "currencies": [
        {
          "currency": "USD",
          "exchange": 4.5
        },
        {
          "currency": "EUR",
          "exchange": 5.5
        }
      ]
    },
    {
      "year": 2025,
      "inflation": 5.2,
      "currencies": [
        {
          "currency": "USD",
          "exchange": 5
        },
        {
          "currency": "EUR",
          "exchange": 6
        }
      ]
    },
    {
      "year": 2026,
      "inflation": 5.26,
      "currencies": [
        {
          "currency": "USD",
          "exchange": 5.55
        },
        {
          "currency": "EUR",
          "exchange": 6.66
        }
      ]
    },
    {
      "year": 2027,
      "inflation": 5.3,
      "currencies": [
        {
          "currency": "USD",
          "exchange": 5.77
        },
        {
          "currency": "EUR",
          "exchange": 6.88
        }
      ]
    }
  ],
  "created_at": "2024-09-24T00:54:59Z",
  "updated_at": null
}
```
### createPlanFC03
```json
HTTP/1.1 201 Created
Content-Type: application/json
Date: Tue, 24 Sep 2024 00:55:26 GMT
Content-Length: 618

{
  "plan_id": "9b19bf6e-bec5-49e7-9f35-06f5ba3de62c",
  "code": "FC 03 2025",
  "name": "Forecast 03 2025",
  "assumptions": [
    {
      "year": 2024,
      "inflation": 4,
      "currencies": [
        {
          "currency": "USD",
          "exchange": 4.5
        },
        {
          "currency": "EUR",
          "exchange": 5.5
        }
      ]
    },
    {
      "year": 2025,
      "inflation": 5.2,
      "currencies": [
        {
          "currency": "USD",
          "exchange": 5
        },
        {
          "currency": "EUR",
          "exchange": 6
        }
      ]
    },
    {
      "year": 2026,
      "inflation": 5.26,
      "currencies": [
        {
          "currency": "USD",
          "exchange": 5.55
        },
        {
          "currency": "EUR",
          "exchange": 6.66
        }
      ]
    },
    {
      "year": 2027,
      "inflation": 5.3,
      "currencies": [
        {
          "currency": "USD",
          "exchange": 5.77
        },
        {
          "currency": "EUR",
          "exchange": 6.88
        }
      ]
    }
  ],
  "created_at": "2024-09-24T00:55:26Z",
  "updated_at": null
}
```
### listPlans
```json
HTTP/1.1 200 OK
Content-Type: application/json
Date: Tue, 24 Sep 2024 00:55:49 GMT
Content-Length: 1247

{
  "plans": [
    {
      "plan_id": "921350d4-b9ce-443e-b284-b0287213919b",
      "code": "BP 2025",
      "name": "Business Plan 2025",
      "assumptions": [
        {
          "year": 2024,
          "inflation": 4,
          "currencies": [
            {
              "currency": "USD",
              "exchange": 4.5
            },
            {
              "currency": "EUR",
              "exchange": 5.5
            }
          ]
        },
        {
          "year": 2025,
          "inflation": 5.2,
          "currencies": [
            {
              "currency": "USD",
              "exchange": 5
            },
            {
              "currency": "EUR",
              "exchange": 6
            }
          ]
        },
        {
          "year": 2026,
          "inflation": 5.26,
          "currencies": [
            {
              "currency": "USD",
              "exchange": 5.55
            },
            {
              "currency": "EUR",
              "exchange": 6.66
            }
          ]
        },
        {
          "year": 2027,
          "inflation": 5.3,
          "currencies": [
            {
              "currency": "USD",
              "exchange": 5.77
            },
            {
              "currency": "EUR",
              "exchange": 6.88
            }
          ]
        }
      ],
      "created_at": "2024-09-24T00:54:59Z",
      "updated_at": null
    },
    {
      "plan_id": "9b19bf6e-bec5-49e7-9f35-06f5ba3de62c",
      "code": "FC 03 2025",
      "name": "Forecast 03 2025",
      "assumptions": [
        {
          "year": 2024,
          "inflation": 4,
          "currencies": [
            {
              "currency": "USD",
              "exchange": 4.5
            },
            {
              "currency": "EUR",
              "exchange": 5.5
            }
          ]
        },
        {
          "year": 2025,
          "inflation": 5.2,
          "currencies": [
            {
              "currency": "USD",
              "exchange": 5
            },
            {
              "currency": "EUR",
              "exchange": 6
            }
          ]
        },
        {
          "year": 2026,
          "inflation": 5.26,
          "currencies": [
            {
              "currency": "USD",
              "exchange": 5.55
            },
            {
              "currency": "EUR",
              "exchange": 6.66
            }
          ]
        },
        {
          "year": 2027,
          "inflation": 5.3,
          "currencies": [
            {
              "currency": "USD",
              "exchange": 5.77
            },
            {
              "currency": "EUR",
              "exchange": 6.88
            }
          ]
        }
      ],
      "created_at": "2024-09-24T00:55:26Z",
      "updated_at": null
    }
  ]
}
```
### updatePlan
```json
HTTP/1.1 200 OK
Content-Type: application/json
Date: Tue, 24 Sep 2024 00:56:12 GMT
Content-Length: 637

{
  "plan_id": "921350d4-b9ce-443e-b284-b0287213919b",
  "code": "BP 2025",
  "name": "Business Plan 2025",
  "assumptions": [
    {
      "year": 2024,
      "inflation": 4,
      "currencies": [
        {
          "currency": "USD",
          "exchange": 4.55
        },
        {
          "currency": "EUR",
          "exchange": 5.55
        }
      ]
    },
    {
      "year": 2025,
      "inflation": 5.2,
      "currencies": [
        {
          "currency": "USD",
          "exchange": 5
        },
        {
          "currency": "EUR",
          "exchange": 6
        }
      ]
    },
    {
      "year": 2026,
      "inflation": 5.26,
      "currencies": [
        {
          "currency": "USD",
          "exchange": 5.55
        },
        {
          "currency": "EUR",
          "exchange": 6.66
        }
      ]
    },
    {
      "year": 2027,
      "inflation": 5.3,
      "currencies": [
        {
          "currency": "USD",
          "exchange": 5.77
        },
        {
          "currency": "EUR",
          "exchange": 6.88
        }
      ]
    }
  ],
  "created_at": "2024-09-24T00:54:59Z",
  "updated_at": "2024-09-24T00:56:12Z"
}
```

### createPortfolioBP
```json
HTTP/1.1 201 Created
Content-Type: application/json
Date: Tue, 24 Sep 2024 00:56:39 GMT
Content-Length: 56

{
  "portfolio_id": "9296d5c2-8a83-4999-89d1-50c958ac0607"
}
```

### createPortfolioFC03
```json
HTTP/1.1 201 Created
Content-Type: application/json
Date: Tue, 24 Sep 2024 00:56:57 GMT
Content-Length: 56

{
  "portfolio_id": "a5eb2598-05ed-4107-9617-612d7cbceb16"
}
```

### listPortfoliosFC03
```json
HTTP/1.1 200 OK
Content-Type: application/json
Date: Tue, 24 Sep 2024 00:58:28 GMT
Content-Length: 418

{
  "portfolios": [
    {
      "portfolio_id": "a5eb2598-05ed-4107-9617-612d7cbceb16",
      "code": "RIT123456789",
      "review": 1,
      "plan_code": "FC 03 2025",
      "title": "Logistics Cost \u0026 Time Management",
      "description": "This project will streamline our internal processes and increase overall efficiency",
      "duration": 12,
      "manager": "John Doe",
      "estimator": "Marie Doe",
      "start_date": "2025-07-01",
      "created_at": "2024-09-24T00:56:57Z",
      "updated_at": null
    }
  ]
}
```
### listPortfoliosBP
```json
HTTP/1.1 200 OK
Content-Type: application/json
Date: Tue, 24 Sep 2024 00:58:56 GMT
Content-Length: 415

{
  "portfolios": [
    {
      "portfolio_id": "9296d5c2-8a83-4999-89d1-50c958ac0607",
      "code": "RIT123456789",
      "review": 1,
      "plan_code": "BP 2025",
      "title": "Logistics Cost \u0026 Time Management",
      "description": "This project will streamline our internal processes and increase overall efficiency",
      "duration": 12,
      "manager": "John Doe",
      "estimator": "Marie Doe",
      "start_date": "2024-12-01",
      "created_at": "2024-09-24T00:56:39Z",
      "updated_at": null
    }
  ]
}
```

### getPortfolioBP
```json
HTTP/1.1 200 OK
Content-Type: application/json
Date: Tue, 24 Sep 2024 00:59:17 GMT
Transfer-Encoding: chunked

{
  "portfolio_id": "9296d5c2-8a83-4999-89d1-50c958ac0607",
  "code": "RIT123456789",
  "review": 1,
  "plan_code": "BP 2025",
  "title": "Logistics Cost \u0026 Time Management",
  "description": "This project will streamline our internal processes and increase overall efficiency",
  "duration": 12,
  "manager": "John Doe",
  "estimator": "Marie Doe",
  "budgets": [
    {
      "budget_id": "23038ac1-11f6-4a33-b5fe-9aebd0c9d83e",
      "portfolio_id": "9296d5c2-8a83-4999-89d1-50c958ac0607",
      "cost_type": "one_time",
      "description": "External Consulting",
      "comment": "estimativa de consultoria externa atualizada",
      "cost_amount": 80000,
      "cost_currency": "USD",
      "cost_tax": 23,
      "cost_apply_inflation": false,
      "amount": 492000,
      "budget_allocations": [
        {
          "year": 2025,
          "month": 6,
          "amount": 307500
        },
        {
          "year": 2025,
          "month": 7,
          "amount": 184500
        }
      ],
      "budget_yearly": [
        {
          "year": 2025,
          "amount": 492000
        }
      ],
      "created_at": "2024-09-24T00:56:39Z",
      "updated_at": null
    },
    {
      "budget_id": "0606e448-5312-469a-ac8b-d1de46136fe8",
      "portfolio_id": "9296d5c2-8a83-4999-89d1-50c958ac0607",
      "cost_type": "one_time",
      "description": "Mão de obra do PO",
      "comment": "estimativa do PO",
      "cost_amount": 180000,
      "cost_currency": "BRL",
      "cost_tax": 0,
      "cost_apply_inflation": true,
      "amount": 187800,
      "budget_allocations": [
        {
          "year": 2024,
          "month": 12,
          "amount": 30000
        },
        {
          "year": 2025,
          "month": 1,
          "amount": 31560
        },
        {
          "year": 2025,
          "month": 2,
          "amount": 31560
        },
        {
          "year": 2025,
          "month": 3,
          "amount": 31560
        },
        {
          "year": 2025,
          "month": 4,
          "amount": 31560
        },
        {
          "year": 2025,
          "month": 5,
          "amount": 31560
        }
      ],
      "budget_yearly": [
        {
          "year": 2024,
          "amount": 30000
        },
        {
          "year": 2025,
          "amount": 157800
        }
      ],
      "created_at": "2024-09-24T00:56:39Z",
      "updated_at": null
    }
  ],
  "workloads": [
    {
      "workload_id": "c67a8ba9-7357-4588-8e82-031dbdbb77a3",
      "portfolio_id": "9296d5c2-8a83-4999-89d1-50c958ac0607",
      "competence_code": "Tech Doc",
      "competence_name": "Technical Documentation - Updated",
      "comment": "considerado Simone e Regina na atividade",
      "hours": 160,
      "workload_allocations": [
        {
          "year": 2024,
          "month": 12,
          "hours": 20
        },
        {
          "year": 2025,
          "month": 1,
          "hours": 20
        },
        {
          "year": 2025,
          "month": 2,
          "hours": 20
        },
        {
          "year": 2025,
          "month": 3,
          "hours": 20
        },
        {
          "year": 2025,
          "month": 4,
          "hours": 20
        },
        {
          "year": 2025,
          "month": 5,
          "hours": 20
        },
        {
          "year": 2025,
          "month": 6,
          "hours": 20
        },
        {
          "year": 2025,
          "month": 7,
          "hours": 20
        }
      ],
      "workload_yearly": [
        {
          "year": 2024,
          "hours": 20
        },
        {
          "year": 2025,
          "hours": 140
        }
      ],
      "created_at": "2024-09-24T00:56:39Z",
      "updated_at": null
    }
  ],
  "start_date": "2024-12-01",
  "created_at": "2024-09-24T00:56:39Z",
  "updated_at": null
}
```
### getPortfolioFC03
```json
HTTP/1.1 200 OK
Content-Type: application/json
Date: Tue, 24 Sep 2024 01:00:10 GMT
Transfer-Encoding: chunked

{
  "portfolio_id": "a5eb2598-05ed-4107-9617-612d7cbceb16",
  "code": "RIT123456789",
  "review": 1,
  "plan_code": "FC 03 2025",
  "title": "Logistics Cost \u0026 Time Management",
  "description": "This project will streamline our internal processes and increase overall efficiency",
  "duration": 12,
  "manager": "John Doe",
  "estimator": "Marie Doe",
  "budgets": [
    {
      "budget_id": "1ee2bd6b-fe57-433f-b992-9a340deda238",
      "portfolio_id": "a5eb2598-05ed-4107-9617-612d7cbceb16",
      "cost_type": "one_time",
      "description": "External Consulting",
      "comment": "estimativa de consultoria externa atualizada",
      "cost_amount": 80000,
      "cost_currency": "USD",
      "cost_tax": 23,
      "cost_apply_inflation": false,
      "amount": 546120,
      "budget_allocations": [
        {
          "year": 2026,
          "month": 1,
          "amount": 341325
        },
        {
          "year": 2026,
          "month": 2,
          "amount": 204795
        }
      ],
      "budget_yearly": [
        {
          "year": 2026,
          "amount": 546120
        }
      ],
      "created_at": "2024-09-24T00:56:57Z",
      "updated_at": null
    },
    {
      "budget_id": "a0f9e389-5ca1-407d-a59f-29ccda960768",
      "portfolio_id": "a5eb2598-05ed-4107-9617-612d7cbceb16",
      "cost_type": "one_time",
      "description": "Mão de obra do PO",
      "comment": "estimativa do PO",
      "cost_amount": 180000,
      "cost_currency": "BRL",
      "cost_tax": 0,
      "cost_apply_inflation": true,
      "amount": 189360,
      "budget_allocations": [
        {
          "year": 2025,
          "month": 7,
          "amount": 31560
        },
        {
          "year": 2025,
          "month": 8,
          "amount": 31560
        },
        {
          "year": 2025,
          "month": 9,
          "amount": 31560
        },
        {
          "year": 2025,
          "month": 10,
          "amount": 31560
        },
        {
          "year": 2025,
          "month": 11,
          "amount": 31560
        },
        {
          "year": 2025,
          "month": 12,
          "amount": 31560
        }
      ],
      "budget_yearly": [
        {
          "year": 2025,
          "amount": 189360
        }
      ],
      "created_at": "2024-09-24T00:56:57Z",
      "updated_at": null
    }
  ],
  "workloads": [
    {
      "workload_id": "c34ca9c0-134e-4d1c-8a91-d35a81005df8",
      "portfolio_id": "a5eb2598-05ed-4107-9617-612d7cbceb16",
      "competence_code": "Tech Doc",
      "competence_name": "Technical Documentation - Updated",
      "comment": "considerado Simone e Regina na atividade",
      "hours": 160,
      "workload_allocations": [
        {
          "year": 2025,
          "month": 7,
          "hours": 20
        },
        {
          "year": 2025,
          "month": 8,
          "hours": 20
        },
        {
          "year": 2025,
          "month": 9,
          "hours": 20
        },
        {
          "year": 2025,
          "month": 10,
          "hours": 20
        },
        {
          "year": 2025,
          "month": 11,
          "hours": 20
        },
        {
          "year": 2025,
          "month": 12,
          "hours": 20
        },
        {
          "year": 2026,
          "month": 1,
          "hours": 20
        },
        {
          "year": 2026,
          "month": 2,
          "hours": 20
        }
      ],
      "workload_yearly": [
        {
          "year": 2025,
          "hours": 120
        },
        {
          "year": 2026,
          "hours": 40
        }
      ],
      "created_at": "2024-09-24T00:56:57Z",
      "updated_at": null
    }
  ],
  "start_date": "2025-07-01",
  "created_at": "2024-09-24T00:56:57Z",
  "updated_at": null
}
```
### deleteCostConsulting
```json
HTTP/1.1 409 Conflict
Content-Type: application/json
Date: Tue, 24 Sep 2024 01:00:34 GMT
Content-Length: 116

{
  "status_code": 409,
  "error": "Conflict",
  "message": "baseline 1941a226-4f1f-450f-8831-ea3e74b480c1 has 2 portfolio(s)"
}
```
### updateCostConsulting
```json
HTTP/1.1 409 Conflict
Content-Type: application/json
Date: Tue, 24 Sep 2024 01:00:53 GMT
Content-Length: 116

{
  "status_code": 409,
  "error": "Conflict",
  "message": "baseline 1941a226-4f1f-450f-8831-ea3e74b480c1 has 2 portfolio(s)"
}
```

### updateBaseline
```json
HTTP/1.1 409 Conflict
Content-Type: application/json
Date: Tue, 24 Sep 2024 01:01:08 GMT
Content-Length: 116

{
  "status_code": 409,
  "error": "Conflict",
  "message": "baseline 1941a226-4f1f-450f-8831-ea3e74b480c1 has 2 portfolio(s)"
}
```
### deletePortfolioBP
```json
HTTP/1.1 204 No Content
Content-Type: application/json
Date: Tue, 24 Sep 2024 01:01:23 GMT
```

### deletePortfolioFC03
```json
HTTP/1.1 204 No Content
Content-Type: application/json
Date: Tue, 24 Sep 2024 01:01:51 GMT
```
