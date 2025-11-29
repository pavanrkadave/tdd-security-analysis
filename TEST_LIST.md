# Test List - Inside-Out TDD

## Phase 1: Value Objects (Simple)

### Threat
- [x] Threat - can be created with name
- [x] Threat - Name() returns the threat name

### Asset
- [x] Asset - can be created with name
- [x] Asset - Name() returns the asset name

### Measure
- [x] Measure - can be created with description
- [x] Measure - Description() returns the measure description

## Phase 2: Enum-like Value Objects

### Likelihood
- [ ] Likelihood - RARE exists
- [ ] Likelihood - OCCASIONAL exists
- [ ] Likelihood - FREQUENT exists
- [ ] Likelihood - String() returns correct name for RARE
- [ ] Likelihood - String() returns correct name for OCCASIONAL
- [ ] Likelihood - String() returns correct name for FREQUENT

### Impact
- [ ] Impact - MARGINAL exists
- [ ] Impact - MODERATE exists
- [ ] Impact - CRITICAL exists
- [ ] Impact - String() returns correct name for MARGINAL
- [ ] Impact - String() returns correct name for MODERATE
- [ ] Impact - String() returns correct name for CRITICAL

### SecurityGoal
- [ ] SecurityGoal - CONFIDENTIALITY exists
- [ ] SecurityGoal - INTEGRITY exists
- [ ] SecurityGoal - AVAILABILITY exists
- [ ] SecurityGoal - String() returns correct name for CONFIDENTIALITY
- [ ] SecurityGoal - String() returns correct name for INTEGRITY
- [ ] SecurityGoal - String() returns correct name for AVAILABILITY

### RiskLevel
- [ ] RiskLevel - LOW exists
- [ ] RiskLevel - MEDIUM exists
- [ ] RiskLevel - HIGH exists
- [ ] RiskLevel - String() returns correct name for LOW
- [ ] RiskLevel - String() returns correct name for MEDIUM
- [ ] RiskLevel - String() returns correct name for HIGH

### TreatmentStatus
- [ ] TreatmentStatus - UNDEFINED exists
- [ ] TreatmentStatus - ACCEPTED exists
- [ ] TreatmentStatus - MITIGATED exists
- [ ] TreatmentStatus - MONITORED exists
- [ ] TreatmentStatus - TRANSFERRED exists
- [ ] TreatmentStatus - String() returns correct names

## Phase 3: Risk Level Calculation (Business Logic Matrix)

### RARE combinations
- [ ] CalculateRiskLevel - RARE + MARGINAL = LOW
- [ ] CalculateRiskLevel - RARE + MODERATE = MEDIUM
- [ ] CalculateRiskLevel - RARE + CRITICAL = HIGH

### OCCASIONAL combinations
- [ ] CalculateRiskLevel - OCCASIONAL + MARGINAL = LOW
- [ ] CalculateRiskLevel - OCCASIONAL + MODERATE = MEDIUM
- [ ] CalculateRiskLevel - OCCASIONAL + CRITICAL = HIGH

### FREQUENT combinations
- [ ] CalculateRiskLevel - FREQUENT + MARGINAL = MEDIUM
- [ ] CalculateRiskLevel - FREQUENT + MODERATE = HIGH
- [ ] CalculateRiskLevel - FREQUENT + CRITICAL = HIGH

## Phase 4: Risk Aggregate Root (Core Entity)

### Risk Creation
- [ ] Risk - can be created with threat, asset, likelihood, impact, security goal
- [ ] Risk - automatically calculates risk level on creation
- [ ] Risk - has a unique ID (UUID)
- [ ] Risk - initial treatment status is UNDEFINED
- [ ] Risk - measure is nil on creation
- [ ] Risk - ID() returns the risk ID
- [ ] Risk - Threat() returns the threat
- [ ] Risk - Asset() returns the asset
- [ ] Risk - SecurityGoal() returns the security goal
- [ ] Risk - Likelihood() returns the likelihood
- [ ] Risk - Impact() returns the impact
- [ ] Risk - RiskLevel() returns the calculated risk level
- [ ] Risk - TreatmentStatus() returns current treatment status
- [ ] Risk - Measure() returns current measure (or nil)

### Risk - Changing Treatment Status
- [ ] Risk - can change treatment status to MITIGATED
- [ ] Risk - can change treatment status to MONITORED
- [ ] Risk - can change treatment status to TRANSFERRED

### Risk - Accepting Risks (Business Rule: Only LOW can be ACCEPTED)
- [ ] Risk - LOW risk can be accepted
- [ ] Risk - MEDIUM risk cannot be accepted (returns error)
- [ ] Risk - HIGH risk cannot be accepted (returns error)
- [ ] Risk - error message explains why non-LOW risk cannot be accepted

### Risk - Measures (Business Rule: ACCEPTED risks cannot have measures)
- [ ] Risk - can add measure when status is UNDEFINED
- [ ] Risk - can add measure when status is MITIGATED
- [ ] Risk - can add measure when status is MONITORED
- [ ] Risk - can add measure when status is TRANSFERRED
- [ ] Risk - cannot add measure when status is ACCEPTED (returns error)
- [ ] Risk - error message explains why accepted risk cannot have measure

### Risk - Accept and Measure Interaction
- [ ] Risk - accepting risk removes existing measure
- [ ] Risk - can add measure after un-accepting (changing from ACCEPTED to another status)

## Phase 5: Repository Interface & Implementation

### Repository Interface (in domain layer)
- [ ] RiskRepository interface - defines FindByID method
- [ ] RiskRepository interface - defines Save method
- [ ] RiskRepository interface - defines Delete method
- [ ] RiskRepository interface - defines FindAll method

### InMemory Repository Implementation
- [ ] InMemoryRiskRepository - can be created
- [ ] InMemoryRiskRepository - Save stores a risk
- [ ] InMemoryRiskRepository - FindByID returns saved risk
- [ ] InMemoryRiskRepository - FindByID returns error when risk not found
- [ ] InMemoryRiskRepository - Save updates existing risk (same ID)
- [ ] InMemoryRiskRepository - FindAll returns empty list when no risks
- [ ] InMemoryRiskRepository - FindAll returns all saved risks
- [ ] InMemoryRiskRepository - Delete removes a risk
- [ ] InMemoryRiskRepository - Delete returns error when risk not found
- [ ] InMemoryRiskRepository - is thread-safe (concurrent Save calls)
- [ ] InMemoryRiskRepository - is thread-safe (concurrent Read/Write)

## Phase 6: Application Service Layer

### Command DTOs
- [ ] CreateRiskCommand - has all required fields as strings
- [ ] AssignMeasureCommand - has risk ID and measure description
- [ ] ChangeTreatmentStatusCommand - has risk ID and new status string

### RiskApplicationService - Create Risk
- [ ] Service - can be created with repository
- [ ] Service - CreateRisk saves new risk to repository
- [ ] Service - CreateRisk returns risk ID
- [ ] Service - CreateRisk converts string fields to value objects
- [ ] Service - CreateRisk handles invalid likelihood string
- [ ] Service - CreateRisk handles invalid impact string
- [ ] Service - CreateRisk handles invalid security goal string

### RiskApplicationService - Assign Measure
- [ ] Service - AssignMeasure loads risk from repository
- [ ] Service - AssignMeasure adds measure to risk
- [ ] Service - AssignMeasure saves updated risk
- [ ] Service - AssignMeasure returns error when risk not found
- [ ] Service - AssignMeasure returns error when risk is accepted

### RiskApplicationService - Change Treatment Status
- [ ] Service - ChangeTreatmentStatus loads risk from repository
- [ ] Service - ChangeTreatmentStatus changes status to MITIGATED
- [ ] Service - ChangeTreatmentStatus changes status to ACCEPTED (if LOW)
- [ ] Service - ChangeTreatmentStatus saves updated risk
- [ ] Service - ChangeTreatmentStatus returns error when risk not found
- [ ] Service - ChangeTreatmentStatus returns error when accepting non-LOW risk
- [ ] Service - ChangeTreatmentStatus handles invalid status string

### RiskApplicationService - Delete Risk
- [ ] Service - DeleteRisk removes risk from repository
- [ ] Service - DeleteRisk returns error when risk not found

### RiskApplicationService - Get Risk
- [ ] Service - GetRisk returns risk by ID
- [ ] Service - GetRisk returns error when risk not found

### RiskApplicationService - List All Risks
- [ ] Service - ListAllRisks returns all risks from repository
- [ ] Service - ListAllRisks returns empty list when no risks exist

## Phase 7: Integration Tests (Optional but Recommended)

- [ ] Integration - Create LOW risk, accept it, verify in repository
- [ ] Integration - Create risk, add measure, change to MITIGATED, verify in repository
- [ ] Integration - Create HIGH risk, attempt to accept, verify error and no change
- [ ] Integration - Full workflow: Create → Add Measure → Change Status → Delete

## Discovered Tests (Add here as you find edge cases)

- [ ] _Add new tests as you discover them during implementation_

---

## Notes for Test Execution

- Start from top to bottom
- Mark tests as [x] when complete
- Each test should have 2-3 commits: RED → GREEN → (optional REFACTOR)
- If a test reveals a missing prerequisite, add it to the list above current test
- Add discovered tests to the bottom section

## Quick Reference: Which Tests Use Real Objects vs Behavior

**State-based tests (Inside-Out):**
- All Value Object tests
- Risk aggregate tests (test state changes)
- Repository tests (test that data is stored/retrieved correctly)
- Service tests with InMemoryRepository (test end-to-end behavior with real objects)

**No mocking needed** - we use real InMemoryRiskRepository throughout!