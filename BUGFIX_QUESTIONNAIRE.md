# Bug Fix: Questionnaire Form Not Showing Saved Answers

## Problem
When users completed the questionnaire and revisited the form, it showed a blank form instead of their previously saved answers or redirecting them to the next step.

## Root Causes
1. **Database allowed duplicate answers**: The `user_answers` table had no unique constraint on `(tracer_id, question_id)`, allowing multiple answers for the same question in the same tracer session.
2. **Backend always inserted new records**: The `UserAnswer.Create()` method only performed INSERT operations, never UPDATE.
3. **Multiple submissions created duplicates**: Each time users submitted the form, new records were created instead of updating existing ones.
4. **Frontend didn't detect completion**: The form always started from group 1, even when users had already completed all question groups.

## Solution Implemented

### 1. Database Changes
- **Added unique constraint**: `uk_user_answers_tracer_question` on `(tracer_id, question_id)` to prevent duplicate answers
- **Cleaned up existing duplicates**: Removed duplicate records, keeping only the most recent answer for each question
- **Updated init.sql**: Added the unique constraint to the schema for future deployments

### 2. Backend Changes (`backend/internal/model/user_answer.go`)
Changed the INSERT query to use `ON DUPLICATE KEY UPDATE`:
```sql
INSERT INTO user_answers (tracer_id, question_id, answer) VALUES (?, ?, ?)
ON DUPLICATE KEY UPDATE answer = VALUES(answer)
```
This ensures that:
- If an answer doesn't exist, it's inserted
- If an answer already exists for that tracer_id + question_id, it's updated

### 3. Frontend Changes (`frontend/src/routes/KuisionerForm.svelte`)
- **Added completion detection**: On mount, the form now checks if the user has already completed the questionnaire
- **Auto-redirect**: If answers for group 2 questions (4, 5, 6) exist, the user is automatically redirected to the next step (upload ijazah)
- **Resume from correct group**: If only group 1 is completed, the form loads the appropriate next group based on their answer
- **Better logging**: Enhanced console logging to track when existing answers are loaded

## Files Modified
1. `backend/internal/model/user_answer.go` - Updated Create method with UPSERT logic
2. `frontend/src/routes/KuisionerForm.svelte` - Enhanced logging
3. `database/init.sql` - Added unique constraint to schema
4. `database/migration_add_unique_constraint.sql` - Migration file for existing databases

## Testing
After the fix:
1. **New user flow**:
   - User logs in and fills the questionnaire (group 1)
   - User submits and progresses to group 2
   - User completes group 2 and is redirected to upload ijazah ✓

2. **Returning user who completed everything**:
   - User logs in and visits kuisioner-form
   - System detects completion and shows success message ✓
   - User is automatically redirected to upload ijazah ✓

3. **Returning user who only completed group 1**:
   - User logs in and visits kuisioner-form
   - System loads their group 1 answers ✓
   - Form shows the appropriate group 2 questions ✓
   - User can complete the remaining questions

4. **Answer updates**:
   - User can update their answers
   - Updated answers replace old ones (not duplicated) ✓

## Deployment Steps
1. ✅ Applied database migration to add unique constraint
2. ✅ Cleaned up duplicate data
3. ✅ Rebuilt backend Docker image
4. ✅ Rebuilt frontend Docker image
5. ✅ Restarted both services

## Verification
All services are running:
- MySQL: healthy
- Backend: running on ports 8099, 9099
- Frontend: running on port 8080

The application is now live at https://tracert.ttmi.my.id/
