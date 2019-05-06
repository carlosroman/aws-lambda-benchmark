Feature: Check API is up
  As a router
  I want to know if API up
  So that I route traffic to it

  Scenario: Check API endpoint is up
    Given the ELB wants to check the API
    When they check the application status
    Then the status code should be 200
    And the API should return the message "API deployed"
