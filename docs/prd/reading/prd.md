# Reading Module PRD

## 1. Module Goal

The reading module helps users read English articles, understand vocabulary and sentence structures, generate exercises, and turn reading into reusable learning assets.

The module should not only translate text. It should help users understand meaning, structure, context, logic, and test-style comprehension.

## 2. Target Users

- Users who want to improve English reading comprehension.
- Users who need guided article reading.
- Users who want sentence-level explanations.
- Users who want AI-generated reading exercises.
- Users preparing for future exam reading tasks such as IELTS reading.

## 3. Core User Jobs

1. I want to choose suitable reading materials.
2. I want to understand words, phrases, and difficult sentences.
3. I want translations without losing the English reading experience.
4. I want AI to explain the article structure and main ideas.
5. I want to practice questions generated from the article.
6. I want to save vocabulary, sentences, and notes.

## 4. MVP User Flow

1. User enters the reading module.
2. User browses the reading library.
3. User selects an article.
4. User reads in the immersive reading page.
5. User selects words, phrases, or sentences for explanation.
6. User optionally turns on bilingual display.
7. User asks AI to analyze the article.
8. User generates and answers reading questions.
9. System shows answers, explanations, and original text references.
10. System saves vocabulary, notes, and reading performance.
11. System recommends the next article or related writing task.

## 5. Reading Library

The reading library contains article cards.

Article card fields:

- Title
- Topic
- Difficulty
- Estimated reading time
- Word count
- Summary
- Keywords
- Target skills
- Source type

Suggested topics:

- Technology
- Education
- Society
- Business
- Culture
- Science
- Environment
- News reading
- Exam-style reading

## 6. Immersive Reading Page

The reading page should prioritize English text while keeping AI help close.

Core features:

- Paragraph reading
- Sentence-level reading
- Translation under English text
- Bilingual display toggle
- Selected text action menu
- Reading progress
- Article notes
- Vocabulary collection
- AI analysis panel
- Generated question panel

Recommended desktop layout:

- Center: article text
- Right panel: explanation, translation, notes, or questions
- Top controls: display mode, difficulty, AI analysis, question generation

## 7. Text Selection Actions

When user selects a word, phrase, sentence, or paragraph, the product should provide context-aware actions.

### 7.1 Word Actions

- Translate
- Explain word meaning in context
- Show part of speech
- Show pronunciation
- Show example sentence
- Add to vocabulary

### 7.2 Phrase Actions

- Translate phrase
- Explain usage
- Show similar expressions
- Add to expression list

### 7.3 Sentence Actions

- Translate sentence
- Break down grammar
- Explain long sentence structure
- Explain function in paragraph
- Rewrite in simpler English
- Add to saved sentences

### 7.4 Paragraph Actions

- Summarize paragraph
- Explain paragraph role
- Generate questions from paragraph
- Extract useful expressions

## 8. AI Article Analysis

AI can analyze the full article.

Analysis sections:

- Core summary
- Paragraph summaries
- Article structure
- Main claim and supporting evidence
- Important vocabulary
- Difficult sentences
- Useful writing expressions
- Possible speaking discussion questions
- Possible writing prompts

## 9. AI-Generated Questions

AI can generate reading exercises based on the article.

Question types:

- Main idea
- Detail
- Inference
- Vocabulary in context
- Sentence function
- Paragraph matching
- Short answer

MVP should keep question generation generic. IELTS-style question formats can be added later through exam content packs.

## 10. Answer Review

After the user answers generated questions, the system should show:

- Correct or incorrect result
- Correct answer
- Explanation
- Original text reference
- Why the user's answer was wrong
- Related reading skill tag

Example skill tags:

- Main idea recognition
- Detail location
- Inference
- Vocabulary in context
- Long sentence comprehension
- Paragraph logic

## 11. Notes and Vocabulary

Users can save:

- Words
- Phrases
- Sentences
- Article notes
- AI explanations
- Useful writing expressions

Saved content should be available for later review and can be reused by the writing module.

## 12. Required Data Objects

### 12.1 Article

- Article ID
- Title
- Content
- Topic
- Difficulty
- Word count
- Estimated reading time
- Keywords
- Source
- Created time

### 12.2 Reading Session

- Session ID
- User ID
- Article ID
- Start time
- End time
- Reading progress
- Selected text actions
- Generated questions
- Answers
- Saved notes
- Saved vocabulary
- AI analysis

### 12.3 Reading Question

- Question ID
- Article ID
- Type
- Question text
- Options
- Correct answer
- Explanation
- Source text reference
- Skill tag

## 13. AI Capability Requirements

The module needs AI capabilities for:

- Translation
- Contextual word explanation
- Long sentence parsing
- Article summarization
- Structure analysis
- Exercise generation
- Answer explanation
- Reading weakness diagnosis
- Cross-module prompt generation for speaking and writing

## 14. MVP Requirements

### 14.1 Must Have

- Reading library
- Article reading page
- Bilingual display toggle
- Text selection actions
- AI full-article analysis
- AI-generated reading questions
- Answer review with explanation
- Vocabulary and sentence saving
- Reading session record

### 14.2 Should Have

- Reading difficulty recommendation
- Related writing prompt generation
- Related speaking discussion question generation
- Article completion report

### 14.3 Not in MVP

- User-uploaded PDF parsing
- Web article import
- Full IELTS reading question bank
- Advanced content management backend
- Collaborative reading

## 15. Success Metrics

- Number of articles completed
- Average reading duration
- Number of selected text explanations
- Number of generated questions answered
- Question accuracy rate
- Number of saved words and sentences
- User rating for AI explanations
