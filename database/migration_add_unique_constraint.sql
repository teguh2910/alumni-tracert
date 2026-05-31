-- Add unique constraint to prevent duplicate answers for same question in same tracer
-- This allows ON DUPLICATE KEY UPDATE to work properly

ALTER TABLE `user_answers`
ADD UNIQUE KEY `uk_user_answers_tracer_question` (`tracer_id`, `question_id`);
