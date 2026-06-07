# AI English Learning Platform PRD

## 1. Product Positioning

This product is a macOS desktop English learning platform. It is built as a modular learning workspace, with three primary training modules:

- Listening and speaking
- Reading
- Writing

The platform is not limited to IELTS. IELTS, TOEFL, CET, business English, interview English, and other exam or scenario packages can be added later as content packs or question banks. The first version focuses on a real, usable MVP for English learners who want structured AI-assisted practice.

The main product structure is a module-based toolbox. AI coaching is embedded across the modules instead of being the only entry point.

## 2. Target Users

### 2.1 Primary Users

- Users preparing for English exams such as IELTS in the future.
- Users who need to improve English listening, speaking, reading, and writing.
- Users who want AI feedback but still prefer a structured learning product over a generic chatbot.
- Users who study on macOS and need a focused desktop learning workspace.

### 2.2 User Needs

- Practice English in realistic scenarios.
- Understand reading materials deeply instead of only translating them.
- Get structured writing feedback and improve through revision.
- Track personal weaknesses and learning progress.
- Receive AI help inside each learning task without losing control of the workflow.

## 3. Product Principles

1. The product is a learning system, not a chat wrapper.
2. Every AI interaction should produce reusable learning value.
3. Each module should have an independent complete workflow.
4. AI feedback should be structured, reviewable, and connected to user history.
5. IELTS and other exams should be supported through future content packs, not hard-coded into the first product boundary.

## 4. Module Structure

```text
prd/
  overview.md
  listening-speaking/
    prd.md
  reading/
    prd.md
  writing/
    prd.md
```

## 5. Core Product Modules

### 5.1 Listening and Speaking

The listening and speaking module helps users train spoken interaction, listening comprehension, pronunciation, fluency, and oral summarization.

Core workflows:

- AI scenario conversation
- Role-play practice
- Shadowing and pronunciation feedback
- Listening retelling
- In-session speaking support
- Post-session AI review
- Weakness tracking

### 5.2 Reading

The reading module helps users read articles, understand difficult language, generate exercises, and turn reading materials into long-term learning assets.

Core workflows:

- Reading library
- Immersive reading page
- Word, phrase, and sentence explanation
- Bilingual translation display
- AI article analysis
- AI-generated questions
- Reading answer review
- Notes and vocabulary collection

### 5.3 Writing

The writing module helps users complete writing tasks through pre-writing planning, drafting, AI review, revision, and expression accumulation.

Core workflows:

- Writing task library
- Pre-writing idea support
- Writing editor
- In-writing AI assistance
- AI scoring and review
- Paragraph-level feedback
- Revision workflow
- Writing review report

## 6. Platform-Level Capabilities

### 6.1 User Learning Profile

The platform records:

- Learning goal
- Current level
- Module preferences
- Training history
- Common mistakes
- Saved words, sentences, and expressions
- Ability tags
- AI-generated next-step suggestions

The learning profile is not a separate heavy workflow in MVP. It works in the background and supports personalized recommendations inside each module.

### 6.2 AI Coaching Layer

AI is embedded into each module with different roles:

- Listening and speaking: conversation partner, examiner, pronunciation coach, review coach
- Reading: translator, grammar explainer, article analyst, question generator
- Writing: brainstorming assistant, reviewer, scoring teacher, revision coach

AI output should be structured whenever possible. For example, a writing review should contain scores, issues, examples, suggested rewrites, and next actions instead of a single long paragraph.

### 6.3 Content and Task System

MVP content types:

- Speaking scenarios
- Role-play prompts
- Listening materials
- Reading articles
- Reading questions
- Writing prompts
- AI-generated exercises

Future content types:

- IELTS reading question banks
- IELTS speaking Part 1, Part 2, and Part 3 packs
- IELTS writing Task 1 and Task 2 packs
- TOEFL, CET, business English, and interview English packs

### 6.4 Learning Records and Review

Each completed training session should generate a structured record:

- Module type
- Content or task
- User input
- AI feedback
- Mistakes
- Suggested expressions
- Ability tag updates
- Recommended next task
- Timestamp and duration

These records support progress tracking, review, and future personalized recommendations.

## 7. macOS Experience

The first version should feel like a desktop learning workspace.

Recommended layout:

- Left sidebar: module navigation and history
- Center area: main training workspace
- Right panel: AI feedback, notes, vocabulary, or review
- Top area: current task, timer, mode controls, and status

The app should support long-form reading, focused writing, and voice-based speaking practice. The UI should prioritize clarity, low friction, and repeated daily use.

## 8. MVP Scope

### 8.1 Included in MVP

- macOS desktop application information architecture
- Three primary modules: listening-speaking, reading, writing
- One complete core workflow for each module
- AI calls and structured AI feedback
- Basic user learning profile
- Training records
- Basic content library
- Reading translation and analysis
- Writing scoring and review
- Speaking conversation and review
- Future question bank extension boundary

### 8.2 Excluded from MVP

- Full IELTS question bank
- Teacher or classroom backend
- Paid subscription system
- Community features
- Human teacher correction
- Large-scale content management backend
- Mobile apps
- Multi-tenant organization management

## 9. Success Metrics

### 9.1 Product Usage

- Daily active users
- Weekly active users
- Average sessions per user
- Average training duration
- Completion rate per module

### 9.2 Learning Engagement

- Number of speaking sessions completed
- Number of reading articles completed
- Number of writing drafts submitted
- Number of revisions completed
- Number of saved words and expressions

### 9.3 AI Feedback Value

- User rating of AI feedback
- Percentage of AI suggestions accepted
- Writing revision improvement rate
- Reading question answer accuracy improvement
- Speaking weakness reduction over time

## 10. Future Expansion

Future versions can add:

- IELTS training packs
- Full question bank system
- Personalized learning plan
- More detailed ability dashboard
- User-uploaded PDF and article parsing
- Web article import
- Spaced review for vocabulary and expressions
- Local-first data storage with cloud sync
- Subscription and paid content packages
