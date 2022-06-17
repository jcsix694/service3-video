# service3-video

## Getting Started
1. Install prerequisites
2. Pull down Repository
3. Run `make all`
4. Run `make kind-up`
5. Run `make kind-load`
6. Run `make kind-apply`


## Endpoints
`GET http://localhost:4000/debug/liveness` - Service Information 
`GET http://localhost:4000/debug/readiness` - Checks if the project is ready
`GET http://localhost:3000/v1/test` - Test Endpoint