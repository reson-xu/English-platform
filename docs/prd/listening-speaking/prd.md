# Listening and Speaking Module PRD

## 1. Module Goal

The listening and speaking module helps users improve spoken English, listening comprehension, pronunciation, fluency, and real-time response ability.

The module should not feel like an open-ended AI chat page. It should feel like a structured training system where AI has clear training roles, session goals, feedback standards, and review outputs.

## 2. Target Users

- Users who want to speak English more fluently.
- Users who need realistic conversation practice.
- Users who want to improve pronunciation and listening response.
- Users preparing for future exam speaking tasks such as IELTS speaking.
- Users who lack a human practice partner.

## 3. Core User Jobs

1. I want to practice English conversation in realistic scenarios.
2. I want AI to challenge me, not just passively reply.
3. I want to know what I said wrong after a conversation.
4. I want help when I cannot express myself.
5. I want to improve pronunciation, fluency, and listening comprehension.
6. I want each practice session to become a trackable learning record.

## 4. MVP User Flow

1. User enters the listening and speaking module.
2. User selects a training mode.
3. User selects a scenario, difficulty, and session duration.
4. AI starts the training session.
5. User speaks or listens based on the mode.
6. AI responds, asks follow-up questions, or gives in-session support.
7. User ends the session.
8. AI generates a structured session review.
9. System saves mistakes, useful expressions, and ability tags.
10. System recommends the next training task.

## 5. Training Modes

### 5.1 AI Scenario Conversation

Users choose a scenario and talk with AI.

Example scenarios:

- Daily conversation
- Travel
- Campus life
- Workplace discussion
- Job interview
- Opinion discussion
- Problem solving
- Future exam speaking practice

AI behavior:

- Adjust vocabulary and sentence complexity based on user level.
- Adjust speaking speed based on user performance.
- Ask follow-up questions.
- Encourage longer answers when user responses are too short.
- Push the user to explain reasons and give examples.

### 5.2 Role-Play Practice

Users choose a role, goal, and conversation context.

Example tasks:

- Persuade a friend to accept your idea.
- Explain a delayed assignment to a teacher.
- Ask hotel staff to solve a room problem.
- Defend your opinion in a group discussion.

AI behavior:

- Play the counterpart role.
- Challenge unclear answers.
- Ask for clarification.
- Create realistic pressure through disagreement or follow-up.
- Evaluate whether the user achieved the communication goal.

### 5.3 Shadowing

AI plays a short sentence, dialogue, or paragraph. User listens and repeats.

Evaluation dimensions:

- Pronunciation accuracy
- Stress
- Intonation
- Pause control
- Rhythm
- Fluency

Output:

- Sentence-level score
- Highlighted weak sounds or rhythm problems
- Suggested reread sentence
- Native-like reference audio or text

### 5.4 Listening Retelling

AI plays a listening material. User retells the content orally.

Evaluation dimensions:

- Key information coverage
- Logical sequence
- Detail retention
- Vocabulary usage
- Grammar accuracy
- Fluency

Output:

- Missing key points
- Incorrect or unclear retelling
- Suggested improved retelling
- Recommended listening difficulty adjustment

### 5.5 In-Session Speaking Support

During a session, user can request help without ending the training.

Support actions:

- Give me an expression suggestion.
- Make this sentence simpler.
- Give me keywords.
- Correct my last sentence.
- Ask the question again.
- Slow down.
- Give me an example answer.

The product should clearly separate support from replacement. AI can guide and suggest, but the user should still complete the speaking task.

## 6. Session Review

After each session, AI generates a structured review.

Review sections:

- Overall performance
- Fluency
- Accuracy
- Vocabulary range
- Grammar issues
- Pronunciation issues
- Listening comprehension
- Communication effectiveness
- Strong sentences
- Sentences to rewrite
- Reusable expressions
- Recommended next practice

## 7. Weakness Tracking

The system should convert repeated issues into user ability tags.

Example tags:

- Past tense usage is unstable.
- Answers lack examples.
- Long sentence organization is weak.
- Pronunciation issue with th sound.
- Pronunciation issue with r and l.
- Retelling misses details.
- User overuses simple connectors.
- User pauses frequently before opinion statements.

Tags are used to recommend later tasks and personalize AI feedback.

## 8. Required Data Objects

### 8.1 Speaking Scenario

- Scenario ID
- Title
- Category
- Difficulty
- User goal
- AI role
- Opening prompt
- Suggested duration
- Target skills

### 8.2 Training Session

- Session ID
- User ID
- Mode
- Scenario ID
- Start time
- End time
- Transcript
- Audio references
- AI feedback
- Scores
- Mistakes
- Ability tag updates

### 8.3 Speaking Feedback

- Overall score
- Fluency score
- Accuracy score
- Vocabulary score
- Pronunciation score
- Listening score
- Key issues
- Improved examples
- Next task suggestion

## 9. AI Capability Requirements

The module needs AI capabilities for:

- Speech recognition
- Text-to-speech
- Dialogue generation
- Role-play behavior
- Pronunciation feedback
- Transcript analysis
- Error detection
- Structured scoring
- Personalized task recommendation

## 10. MVP Requirements

### 10.1 Must Have

- Scenario conversation mode
- Role-play mode
- Basic shadowing mode
- Listening retelling mode
- Audio input and transcript display
- AI voice output
- In-session support actions
- Session review
- Learning record save
- Weakness tag extraction

### 10.2 Should Have

- Session replay
- Sentence-level correction
- Difficulty adjustment
- Saved expressions
- User-selectable AI role

### 10.3 Not in MVP

- Full IELTS speaking question bank
- Human teacher feedback
- Real-time multiplayer speaking
- Advanced phoneme-level visualization
- Video conversation

## 11. Success Metrics

- Number of completed speaking sessions
- Average session duration
- Percentage of sessions with review opened
- Number of repeated practice attempts
- Reduction of repeated weakness tags
- User rating for AI conversation quality
- User rating for pronunciation feedback usefulness
