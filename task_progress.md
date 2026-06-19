# Task Progress — Quiz Revamp

## Todo

- [x] Analyze current codebase structure
- [x] Create 15 new Dark Triad questions with metadata (id, trait, reverse, text)
- [x] Update quiz.html: replace questions, change to 5-point scale, add scoring logic
- [x] Update services/quiz.go: new scoring system with per-trait calculation + normalization
- [x] Update handlers/quiz.go: handle 15 answers instead of 5
- [x] Update models/user.go: keep existing fields (scoring stays per-trait)
- [x] Update repositories/user.go: no changes needed (already stores 3 scores)
- [x] Update hasil.html: professional interpretations instead of negative labels
- [x] Verify all changes compile and work together
