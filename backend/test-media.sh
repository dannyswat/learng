#!/bin/bash

# Test script for Media Upload endpoints (Sprint 2 - Media)
# Tests image and audio upload functionality

BASE_URL="http://localhost:8080/api/v1"
TEMP_DIR="/tmp/learng-media-test"

# Colors for output
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Counters
TESTS_RUN=0
TESTS_PASSED=0
TESTS_FAILED=0

# Create temp directory for test files
mkdir -p "$TEMP_DIR"

# Function to print colored output
print_status() {
    if [ $1 -eq 0 ]; then
        echo -e "${GREEN}✓ PASS${NC}: $2"
        ((TESTS_PASSED++))
    else
        echo -e "${RED}✗ FAIL${NC}: $2"
        ((TESTS_FAILED++))
    fi
    ((TESTS_RUN++))
}

# Function to create test image file (1x1 PNG)
create_test_image() {
    local filename="$1"
    # Create a minimal 1x1 PNG file (base64 encoded)
    echo "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mNk+M9QDwADhgGAWjR9awAAAABJRU5ErkJggg==" | base64 -d > "$filename"
}

# Function to create test audio file (minimal WAV)
create_test_audio() {
    local filename="$1"
    # Create a minimal WAV file header
    printf "RIFF\x24\x00\x00\x00WAVEfmt \x10\x00\x00\x00\x01\x00\x01\x00\x44\xAC\x00\x00\x88\x58\x01\x00\x02\x00\x10\x00data\x00\x00\x00\x00" > "$filename"
}

echo "================================================"
echo "  Media Upload Endpoints Test Suite"
echo "================================================"
echo ""

# Step 1: Register a test user
echo "1. Setting up test user..."
REGISTER_RESPONSE=$(curl -s -X POST "$BASE_URL/auth/register" \
  -H "Content-Type: application/json" \
  -d '{
    "email": "mediatest@example.com",
    "password": "testpass123",
    "displayName": "Media Test User",
    "role": "admin"
  }')

if echo "$REGISTER_RESPONSE" | grep -q '"token"'; then
    print_status 0 "User registration"
    TOKEN=$(echo "$REGISTER_RESPONSE" | grep -o '"token":"[^"]*"' | cut -d'"' -f4)
else
    # Try to login if user already exists
    LOGIN_RESPONSE=$(curl -s -X POST "$BASE_URL/auth/login" \
      -H "Content-Type: application/json" \
      -d '{
        "email": "mediatest@example.com",
        "password": "testpass123"
      }')
    
    if echo "$LOGIN_RESPONSE" | grep -q '"token"'; then
        print_status 0 "User login (existing user)"
        TOKEN=$(echo "$LOGIN_RESPONSE" | grep -o '"token":"[^"]*"' | cut -d'"' -f4)
    else
        print_status 1 "Authentication failed"
        echo "Response: $LOGIN_RESPONSE"
        exit 1
    fi
fi

echo ""

# Test 2: Upload PNG Image
echo "2. Testing image upload (PNG)..."
create_test_image "$TEMP_DIR/test.png"
IMAGE_RESPONSE=$(curl -s -X POST "$BASE_URL/media/upload/image" \
  -H "Authorization: Bearer $TOKEN" \
  -F "file=@$TEMP_DIR/test.png")

if echo "$IMAGE_RESPONSE" | grep -q '"url"'; then
    IMAGE_URL=$(echo "$IMAGE_RESPONSE" | grep -o '"url":"[^"]*"' | cut -d'"' -f4)
    print_status 0 "Image upload (PNG) - URL: $IMAGE_URL"
else
    print_status 1 "Image upload (PNG)"
    echo "Response: $IMAGE_RESPONSE"
fi

echo ""

# Test 3: Upload JPEG Image
echo "3. Testing image upload (JPEG)..."
# Create JPEG test file (minimal JPEG header)
printf "\xFF\xD8\xFF\xE0\x00\x10JFIF\x00\x01\x01\x00\x00\x01\x00\x01\x00\x00\xFF\xD9" > "$TEMP_DIR/test.jpg"
JPEG_RESPONSE=$(curl -s -X POST "$BASE_URL/media/upload/image" \
  -H "Authorization: Bearer $TOKEN" \
  -F "file=@$TEMP_DIR/test.jpg")

if echo "$JPEG_RESPONSE" | grep -q '"url"'; then
    print_status 0 "Image upload (JPEG)"
else
    print_status 1 "Image upload (JPEG)"
    echo "Response: $JPEG_RESPONSE"
fi

echo ""

# Test 4: Upload Audio (WAV)
echo "4. Testing audio upload (WAV)..."
create_test_audio "$TEMP_DIR/test.wav"
AUDIO_RESPONSE=$(curl -s -X POST "$BASE_URL/media/upload/audio" \
  -H "Authorization: Bearer $TOKEN" \
  -F "file=@$TEMP_DIR/test.wav")

if echo "$AUDIO_RESPONSE" | grep -q '"url"'; then
    AUDIO_URL=$(echo "$AUDIO_RESPONSE" | grep -o '"url":"[^"]*"' | cut -d'"' -f4)
    print_status 0 "Audio upload (WAV) - URL: $AUDIO_URL"
else
    print_status 1 "Audio upload (WAV)"
    echo "Response: $AUDIO_RESPONSE"
fi

echo ""

# Test 5: Upload without file (should fail)
echo "5. Testing upload without file (should fail)..."
NO_FILE_RESPONSE=$(curl -s -X POST "$BASE_URL/media/upload/image" \
  -H "Authorization: Bearer $TOKEN")

if echo "$NO_FILE_RESPONSE" | grep -q '"error"'; then
    print_status 0 "Validation: No file uploaded"
else
    print_status 1 "Validation: No file uploaded"
    echo "Response: $NO_FILE_RESPONSE"
fi

echo ""

# Test 6: Upload oversized file (should fail)
echo "6. Testing oversized image upload (should fail)..."
# Create a file larger than 5MB
dd if=/dev/zero of="$TEMP_DIR/large.png" bs=1024 count=6000 2>/dev/null
LARGE_RESPONSE=$(curl -s -X POST "$BASE_URL/media/upload/image" \
  -H "Authorization: Bearer $TOKEN" \
  -F "file=@$TEMP_DIR/large.png")

if echo "$LARGE_RESPONSE" | grep -q '"error"'; then
    print_status 0 "Validation: File too large"
else
    print_status 1 "Validation: File too large"
    echo "Response: $LARGE_RESPONSE"
fi

echo ""

# Test 7: Upload invalid file type (should fail)
echo "7. Testing invalid file type for image (should fail)..."
echo "This is not an image" > "$TEMP_DIR/test.txt"
INVALID_RESPONSE=$(curl -s -X POST "$BASE_URL/media/upload/image" \
  -H "Authorization: Bearer $TOKEN" \
  -F "file=@$TEMP_DIR/test.txt")

if echo "$INVALID_RESPONSE" | grep -q '"error"'; then
    print_status 0 "Validation: Invalid file type"
else
    print_status 1 "Validation: Invalid file type"
    echo "Response: $INVALID_RESPONSE"
fi

echo ""

# Test 8: Upload without authentication (should fail)
echo "8. Testing upload without authentication (should fail)..."
NO_AUTH_RESPONSE=$(curl -s -X POST "$BASE_URL/media/upload/image" \
  -F "file=@$TEMP_DIR/test.png")

if echo "$NO_AUTH_RESPONSE" | grep -q '"error"'; then
    print_status 0 "Authentication required"
else
    print_status 1 "Authentication required"
    echo "Response: $NO_AUTH_RESPONSE"
fi

echo ""

# Test 9: Verify uploaded files are accessible
echo "9. Testing uploaded files are accessible..."
if [ -n "$IMAGE_URL" ]; then
    HTTP_CODE=$(curl -s -o /dev/null -w "%{http_code}" "http://localhost:8080$IMAGE_URL")
    if [ "$HTTP_CODE" -eq 200 ]; then
        print_status 0 "Image file accessible"
    else
        print_status 1 "Image file accessible (HTTP $HTTP_CODE)"
    fi
else
    print_status 1 "Image file accessible (no URL)"
fi

if [ -n "$AUDIO_URL" ]; then
    HTTP_CODE=$(curl -s -o /dev/null -w "%{http_code}" "http://localhost:8080$AUDIO_URL")
    if [ "$HTTP_CODE" -eq 200 ]; then
        print_status 0 "Audio file accessible"
    else
        print_status 1 "Audio file accessible (HTTP $HTTP_CODE)"
    fi
else
    print_status 1 "Audio file accessible (no URL)"
fi

echo ""

# Test 10: Create word with uploaded media
echo "10. Testing word creation with uploaded media..."
# First create a journey
JOURNEY_RESPONSE=$(curl -s -X POST "$BASE_URL/journeys" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Media Test Journey",
    "description": "Journey for testing media uploads",
    "sourceLanguage": "en",
    "targetLanguage": "zh-HK"
  }')

JOURNEY_ID=$(echo "$JOURNEY_RESPONSE" | grep -o '"id":"[^"]*"' | head -1 | cut -d'"' -f4)

# Create a scenario
SCENARIO_RESPONSE=$(curl -s -X POST "$BASE_URL/scenarios" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d "{
    \"journeyId\": \"$JOURNEY_ID\",
    \"title\": \"Media Test Scenario\",
    \"description\": \"Scenario for testing media\"
  }")

SCENARIO_ID=$(echo "$SCENARIO_RESPONSE" | grep -o '"id":"[^"]*"' | head -1 | cut -d'"' -f4)

# Create word with media
if [ -n "$IMAGE_URL" ] && [ -n "$AUDIO_URL" ]; then
    WORD_RESPONSE=$(curl -s -X POST "$BASE_URL/words" \
      -H "Authorization: Bearer $TOKEN" \
      -H "Content-Type: application/json" \
      -d "{
        \"scenarioId\": \"$SCENARIO_ID\",
        \"sourceText\": \"apple\",
        \"targetText\": \"蘋果\",
        \"imageUrl\": \"$IMAGE_URL\",
        \"audioUrl\": \"$AUDIO_URL\",
        \"generationMethod\": \"manual\"
      }")
    
    if echo "$WORD_RESPONSE" | grep -q '"id"'; then
        print_status 0 "Word created with uploaded media"
    else
        print_status 1 "Word created with uploaded media"
        echo "Response: $WORD_RESPONSE"
    fi
else
    print_status 1 "Word created with uploaded media (missing URLs)"
fi

echo ""

# Cleanup
echo "Cleaning up test files..."
rm -rf "$TEMP_DIR"

# Summary
echo "================================================"
echo "  Test Summary"
echo "================================================"
echo -e "Total Tests:  $TESTS_RUN"
echo -e "${GREEN}Passed:       $TESTS_PASSED${NC}"
echo -e "${RED}Failed:       $TESTS_FAILED${NC}"
echo ""

if [ $TESTS_FAILED -eq 0 ]; then
    echo -e "${GREEN}✓ All tests passed!${NC}"
    exit 0
else
    echo -e "${RED}✗ Some tests failed${NC}"
    exit 1
fi
