# Writing Module PRD

## 1. Module Goal

The writing module helps users complete English writing practice through planning, drafting, AI review, revision, and expression accumulation.

The module should not stop at "write and let AI correct it." It should guide users through a full writing improvement loop: understand the task, form ideas, draft, review, revise, compare, and learn reusable expressions.

## 2. Target Users

- Users who want to improve English writing.
- Users who need structured correction and feedback.
- Users who want help developing ideas and arguments.
- Users who want to practice exam-style writing in the future.
- Users who want to turn reading input into writing output.

## 3. Core User Jobs

1. I want to choose or create a writing task.
2. I want help understanding the task and building an outline.
3. I want to write in a focused editor.
4. I want AI to improve my writing without fully replacing my work.
5. I want structured scoring and review.
6. I want to revise and know whether the second draft is better.
7. I want to save useful expressions for future writing.

## 4. MVP User Flow

1. User enters the writing module.
2. User selects a writing task or starts free writing.
3. User uses AI pre-writing support if needed.
4. User writes in the editor.
5. User selects local text for in-writing AI assistance if needed.
6. User submits the draft for AI scoring and review.
7. AI returns structured feedback.
8. User revises the draft.
9. AI compares draft 1 and draft 2.
10. System saves scores, mistakes, expressions, and ability tags.
11. System recommends the next writing task.

## 5. Writing Task Library

Task types:

- Free writing
- Opinion essay
- Chart or data description
- Email or letter
- Summary
- Rewrite
- Reading-based writing
- Exam-style writing

Future exam packs can include IELTS Task 1 and Task 2.

Task fields:

- Task ID
- Title
- Type
- Prompt
- Difficulty
- Suggested word count
- Suggested time
- Target skills
- Source article ID if related to reading

## 6. Pre-Writing Support

Before writing, users can ask AI to help with:

- Understand the prompt
- Identify keywords
- Explain task requirements
- Generate possible positions
- Provide supporting ideas
- Create an outline
- Suggest paragraph structure
- Suggest useful expressions

AI should support thinking, not replace the user's writing.

## 7. Writing Editor

The editor should support focused long-form writing.

Core features:

- Word count
- Timer
- Draft save
- Version history
- Paragraph display
- Submit for review
- Local text selection
- AI assistance panel

Recommended desktop layout:

- Center: writing editor
- Right panel: AI assistance, outline, review, or expression list
- Top controls: task info, timer, word count, submit button

## 8. In-Writing AI Assistance

When user selects a sentence or paragraph, AI can provide:

- Grammar check
- Natural expression rewrite
- More advanced vocabulary option
- Simpler expression option
- Argument expansion
- Logic check
- Next sentence suggestion
- Transition suggestion

The product should avoid making AI the default writer. The UX should frame these actions as help, correction, and inspiration.

## 9. AI Scoring and Review

After submitting a draft, AI should return structured review.

Review dimensions:

- Task response
- Content completeness
- Organization and structure
- Coherence
- Grammar accuracy
- Vocabulary range
- Clarity
- Overall score

For future IELTS packs, the scoring dimensions can map to:

- Task Achievement or Task Response
- Coherence and Cohesion
- Lexical Resource
- Grammatical Range and Accuracy

## 10. Paragraph-Level Feedback

AI should review each paragraph.

Feedback fields:

- Paragraph role
- Main issue
- Logic problem
- Grammar problem
- Vocabulary problem
- Suggested rewrite
- Useful expression
- Strength of the paragraph

## 11. Revision Workflow

The revision workflow is a required part of the learning loop.

Steps:

1. User reads AI review.
2. User edits the draft.
3. User submits revised draft.
4. AI compares original and revised versions.
5. AI explains what improved and what still needs work.
6. System saves both versions and final feedback.

Comparison dimensions:

- Structure improvement
- Grammar improvement
- Vocabulary improvement
- Argument clarity
- Error reduction
- Remaining issues

## 12. Expression Accumulation

Users can save:

- Good sentences from their own drafts
- AI-suggested rewrites
- Useful expressions from reviews
- Expressions imported from reading articles

Saved expressions can be reused in future writing tasks.

## 13. Required Data Objects

### 13.1 Writing Task

- Task ID
- Title
- Type
- Prompt
- Difficulty
- Suggested word count
- Suggested time
- Target skills
- Source article ID

### 13.2 Writing Draft

- Draft ID
- User ID
- Task ID
- Content
- Version number
- Word count
- Created time
- Updated time
- Submit status

### 13.3 Writing Review

- Review ID
- Draft ID
- Overall score
- Dimension scores
- Paragraph feedback
- Grammar issues
- Vocabulary issues
- Structure issues
- Suggested rewrites
- Saved expressions
- Next task suggestion

### 13.4 Revision Comparison

- Comparison ID
- Original draft ID
- Revised draft ID
- Improvements
- Remaining issues
- Score change
- Recommended next revision

## 14. AI Capability Requirements

The module needs AI capabilities for:

- Prompt understanding
- Outline generation
- Grammar correction
- Style rewrite
- Argument improvement
- Structured scoring
- Paragraph-level review
- Draft comparison
- Expression extraction
- Personalized writing recommendations

## 15. MVP Requirements

### 15.1 Must Have

- Writing task library
- Free writing mode
- Writing editor
- Pre-writing AI support
- Local text AI assistance
- AI scoring and review
- Paragraph-level feedback
- Revision workflow
- Draft version record
- Saved expressions

### 15.2 Should Have

- Timer-based writing practice
- Reading-to-writing task generation
- Score trend
- Common mistake summary

### 15.3 Not in MVP

- Full IELTS writing question bank
- Human teacher correction
- Plagiarism detection
- Collaborative writing
- Full classroom workflow

## 16. Success Metrics

- Number of writing tasks started
- Number of drafts submitted
- Revision completion rate
- Average score improvement after revision
- Number of saved expressions
- Number of repeated writing issues
- User rating for AI review quality
