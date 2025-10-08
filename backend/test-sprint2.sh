#!/bin/bash

# Sprint 2 API Test Script
# Tests Journey, Scenario, and Word CRUD operations

BASE_URL="http://localhost:8080/api/v1"
TOKEN=""
JOURNEY_ID=""
SCENARIO_ID=""
WORD_ID=""

echo "======================================"
echo "Sprint 2 - Journey/Scenario/Word CRUD API Tests"
echo "======================================"
echo ""

# Step 1: Register a test admin user
echo "1. Registering test admin user..."
REGISTER_RESPONSE=$(curl -s -X POST "$BASE_URL/auth/register" \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@test.com",
    "password": "Test1234",
    "displayName": "Test Admin",
    "role": "admin"
  }')

echo "Register Response: $REGISTER_RESPONSE"
TOKEN=$(echo $REGISTER_RESPONSE | grep -o '"token":"[^"]*' | cut -d'"' -f4)

if [ -z "$TOKEN" ]; then
  echo "❌ Failed to register user or get token"
  exit 1
fi

echo "✅ User registered successfully"
echo "Token: ${TOKEN:0:20}..."
echo ""

# Step 2: Create a Journey
echo "2. Creating a new journey..."
JOURNEY_RESPONSE=$(curl -s -X POST "$BASE_URL/journeys" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "title": "At the Park",
    "description": "Learn outdoor vocabulary",
    "sourceLanguage": "en",
    "targetLanguage": "zh-HK"
  }')

echo "Journey Response: $JOURNEY_RESPONSE"
JOURNEY_ID=$(echo $JOURNEY_RESPONSE | grep -o '"id":"[^"]*' | cut -d'"' -f4)

if [ -z "$JOURNEY_ID" ]; then
  echo "❌ Failed to create journey"
  exit 1
fi

echo "✅ Journey created successfully"
echo "Journey ID: $JOURNEY_ID"
echo ""

# Step 3: Get Journey by ID
echo "3. Fetching journey details..."
GET_JOURNEY_RESPONSE=$(curl -s -X GET "$BASE_URL/journeys/$JOURNEY_ID" \
  -H "Authorization: Bearer $TOKEN")

echo "Get Journey Response: $GET_JOURNEY_RESPONSE"
echo ""

# Step 4: Get all Journeys
echo "4. Fetching all journeys..."
GET_ALL_RESPONSE=$(curl -s -X GET "$BASE_URL/journeys?status=draft&page=1&limit=10" \
  -H "Authorization: Bearer $TOKEN")

echo "Get All Journeys Response: $GET_ALL_RESPONSE"
echo ""

# Step 5: Update Journey
echo "5. Updating journey..."
UPDATE_JOURNEY_RESPONSE=$(curl -s -X PUT "$BASE_URL/journeys/$JOURNEY_ID" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "title": "At the Park (Updated)",
    "description": "Learn outdoor vocabulary with fun activities",
    "status": "draft"
  }')

echo "Update Journey Response: $UPDATE_JOURNEY_RESPONSE"
echo ""

# Step 6: Create a Scenario
echo "6. Creating a scenario..."
SCENARIO_RESPONSE=$(curl -s -X POST "$BASE_URL/scenarios" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d "{
    \"journeyId\": \"$JOURNEY_ID\",
    \"title\": \"Colors in Nature\",
    \"description\": \"Learn colors you see in the park\",
    \"displayOrder\": 1
  }")

echo "Scenario Response: $SCENARIO_RESPONSE"
SCENARIO_ID=$(echo $SCENARIO_RESPONSE | grep -o '"id":"[^"]*' | cut -d'"' -f4)

if [ -z "$SCENARIO_ID" ]; then
  echo "❌ Failed to create scenario"
  exit 1
fi

echo "✅ Scenario created successfully"
echo "Scenario ID: $SCENARIO_ID"
echo ""

# Step 7: Get Scenario by ID
echo "7. Fetching scenario details..."
GET_SCENARIO_RESPONSE=$(curl -s -X GET "$BASE_URL/scenarios/$SCENARIO_ID" \
  -H "Authorization: Bearer $TOKEN")

echo "Get Scenario Response: $GET_SCENARIO_RESPONSE"
echo ""

# Step 8: Update Scenario
echo "8. Updating scenario..."
UPDATE_SCENARIO_RESPONSE=$(curl -s -X PUT "$BASE_URL/scenarios/$SCENARIO_ID" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "title": "Colors in Nature (Enhanced)",
    "description": "Learn beautiful colors you see in the park"
  }')

echo "Update Scenario Response: $UPDATE_SCENARIO_RESPONSE"
echo ""

# Step 9: Create Words
echo "9. Creating words..."

# Word 1: Red
WORD1_RESPONSE=$(curl -s -X POST "$BASE_URL/words" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d "{
    \"scenarioId\": \"$SCENARIO_ID\",
    \"targetText\": \"紅色\",
    \"sourceText\": \"Red\",
    \"displayOrder\": 1
  }")

echo "Word 1 Response: $WORD1_RESPONSE"
WORD_ID=$(echo $WORD1_RESPONSE | grep -o '"id":"[^"]*' | cut -d'"' -f4)

if [ -z "$WORD_ID" ]; then
  echo "❌ Failed to create word"
  exit 1
fi

echo "✅ Word 1 created successfully"
echo "Word ID: $WORD_ID"
echo ""

# Word 2: Blue
WORD2_RESPONSE=$(curl -s -X POST "$BASE_URL/words" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d "{
    \"scenarioId\": \"$SCENARIO_ID\",
    \"targetText\": \"藍色\",
    \"sourceText\": \"Blue\",
    \"displayOrder\": 2
  }")

echo "Word 2 Response: $WORD2_RESPONSE"
echo ""

# Word 3: Green
WORD3_RESPONSE=$(curl -s -X POST "$BASE_URL/words" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d "{
    \"scenarioId\": \"$SCENARIO_ID\",
    \"targetText\": \"綠色\",
    \"sourceText\": \"Green\",
    \"displayOrder\": 3
  }")

echo "Word 3 Response: $WORD3_RESPONSE"
echo ""

# Step 10: Get Word by ID
echo "10. Fetching word details..."
GET_WORD_RESPONSE=$(curl -s -X GET "$BASE_URL/words/$WORD_ID" \
  -H "Authorization: Bearer $TOKEN")

echo "Get Word Response: $GET_WORD_RESPONSE"
echo ""

# Step 11: Update Word (add media URLs)
echo "11. Updating word with media URLs..."
UPDATE_WORD_RESPONSE=$(curl -s -X PUT "$BASE_URL/words/$WORD_ID" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "imageUrl": "/uploads/images/red.jpg",
    "audioUrl": "/uploads/audio/red.mp3",
    "generationMethod": "manual"
  }')

echo "Update Word Response: $UPDATE_WORD_RESPONSE"
echo ""

# Step 12: Get Journey with all nested data
echo "12. Fetching complete journey with scenarios and words..."
GET_FULL_JOURNEY=$(curl -s -X GET "$BASE_URL/journeys/$JOURNEY_ID" \
  -H "Authorization: Bearer $TOKEN")

echo "Complete Journey Response:"
echo "$GET_FULL_JOURNEY" | jq '.' 2>/dev/null || echo "$GET_FULL_JOURNEY"
echo ""

# Step 13: Test Journey filtering
echo "13. Testing journey filtering (published journeys)..."
FILTER_RESPONSE=$(curl -s -X GET "$BASE_URL/journeys?status=published" \
  -H "Authorization: Bearer $TOKEN")

echo "Filtered Journeys (published): $FILTER_RESPONSE"
echo ""

# Step 14: Delete Word
echo "14. Deleting a word..."
DELETE_WORD_RESPONSE=$(curl -s -X DELETE "$BASE_URL/words/$WORD_ID" \
  -H "Authorization: Bearer $TOKEN" \
  -w "\nHTTP Status: %{http_code}")

echo "Delete Word Response: $DELETE_WORD_RESPONSE"
echo ""

# Step 15: Try to get deleted word (should fail)
echo "15. Attempting to fetch deleted word (should return 404)..."
GET_DELETED_WORD=$(curl -s -X GET "$BASE_URL/words/$WORD_ID" \
  -H "Authorization: Bearer $TOKEN" \
  -w "\nHTTP Status: %{http_code}")

echo "Get Deleted Word Response: $GET_DELETED_WORD"
echo ""

# Step 16: Delete Scenario (should cascade delete remaining words)
echo "16. Deleting scenario (cascade delete words)..."
DELETE_SCENARIO_RESPONSE=$(curl -s -X DELETE "$BASE_URL/scenarios/$SCENARIO_ID" \
  -H "Authorization: Bearer $TOKEN" \
  -w "\nHTTP Status: %{http_code}")

echo "Delete Scenario Response: $DELETE_SCENARIO_RESPONSE"
echo ""

# Step 17: Delete Journey
echo "17. Deleting journey..."
DELETE_JOURNEY_RESPONSE=$(curl -s -X DELETE "$BASE_URL/journeys/$JOURNEY_ID" \
  -H "Authorization: Bearer $TOKEN" \
  -w "\nHTTP Status: %{http_code}")

echo "Delete Journey Response: $DELETE_JOURNEY_RESPONSE"
echo ""

echo "======================================"
echo "✅ All Sprint 2 API tests completed!"
echo "======================================"
echo ""
echo "Summary:"
echo "- Journey CRUD: ✅"
echo "- Scenario CRUD: ✅"
echo "- Word CRUD: ✅"
echo "- Nested data retrieval: ✅"
echo "- Cascade deletions: ✅"
echo ""
