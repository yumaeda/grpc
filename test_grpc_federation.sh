#!/bin/bash

# Color codes for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "${BLUE}=== Testing gRPC Federation Services ===${NC}\n"

# Test Restaurant Service
echo -e "${GREEN}Testing RestaurantService...${NC}"
echo "Request: {\"id\":\"0b95fe0d-a323-43d7-954b-2164a1d3242d\"}"
grpcurl -plaintext -d '{"id":"0b95fe0d-a323-43d7-954b-2164a1d3242d"}' localhost:50051 swapi.restaurant.RestaurantService/GetRestaurant
echo -e "\n"

# Test Video Service
echo -e "${GREEN}Testing VideoService...${NC}"
echo "Request: {\"id\":1}"
grpcurl -plaintext -d '{"id":1}' localhost:50052 swapi.video.VideoService/GetVideo
echo -e "\n"

# Test Photo Service
echo -e "${GREEN}Testing PhotoService...${NC}"
echo "Request: {\"id\":100}"
grpcurl -plaintext -d '{"id":100}' localhost:50053 swapi.photo.PhotoService/GetPhoto
echo -e "\n"

# Test SWAPI Service
echo -e "${GREEN}Testing SWAPIService...${NC}"
echo "Request: {\"id\":\"da04f5c9-ffb0-11ea-ba65-065a10bcba76\"}"
grpcurl -plaintext -d '{"id":"da04f5c9-ffb0-11ea-ba65-065a10bcba76"}' localhost:50054 swapi.SWAPI/GetRestaurant
echo -e "\n"

echo -e "${BLUE}=== All tests completed ===${NC}"
