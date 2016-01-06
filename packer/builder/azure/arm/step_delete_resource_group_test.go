// Copyright (c) Microsoft Open Technologies, Inc.
// All Rights Reserved.
// Licensed under the Apache License, Version 2.0.
// See License.txt in the project root for license information.

package arm

import (
	"fmt"
	"testing"

	"github.com/Azure/packer-azure/packer/builder/azure/common/constants"
	"github.com/mitchellh/multistep"
)

func TestStepDeleteResourceGroupShouldFailIfValidateFails(t *testing.T) {
	var testSubject = &StepDeleteResourceGroup{
		delete: func(string) error { return fmt.Errorf("!! Unit Test FAIL !!") },
		say:    func(message string) {},
		error:  func(e error) {},
	}

	stateBag := DeleteTestStateBagStepDeleteResourceGroup()

	var result = testSubject.Run(stateBag)
	if result != multistep.ActionHalt {
		t.Fatalf("Expected the step to return 'ActionHalt', but got '%s'.", result)
	}

	if _, ok := stateBag.GetOk(constants.Error); ok == false {
		t.Fatalf("Expected the step to set stateBag['%s'], but it was not.", constants.Error)
	}
}

func TestStepDeleteResourceGroupShouldPassIfValidatePasses(t *testing.T) {
	var testSubject = &StepDeleteResourceGroup{
		delete: func(string) error { return nil },
		say:    func(message string) {},
		error:  func(e error) {},
	}

	stateBag := DeleteTestStateBagStepDeleteResourceGroup()

	var result = testSubject.Run(stateBag)
	if result != multistep.ActionContinue {
		t.Fatalf("Expected the step to return 'ActionContinue', but got '%s'.", result)
	}

	if _, ok := stateBag.GetOk(constants.Error); ok == true {
		t.Fatalf("Expected the step to not set stateBag['%s'], but it was.", constants.Error)
	}
}

func TestStepDeleteResourceGroupShouldTakeValidateArgumentsFromStateBag(t *testing.T) {
	var actualResourceGroupName string

	var testSubject = &StepDeleteResourceGroup{
		delete: func(resourceGroupName string) error {
			actualResourceGroupName = resourceGroupName
			return nil
		},
		say:   func(message string) {},
		error: func(e error) {},
	}

	stateBag := DeleteTestStateBagStepDeleteResourceGroup()
	var result = testSubject.Run(stateBag)

	if result != multistep.ActionContinue {
		t.Fatalf("Expected the step to return 'ActionContinue', but got '%s'.", result)
	}

	var expectedResourceGroupName = stateBag.Get(constants.ArmResourceGroupName).(string)

	if actualResourceGroupName != expectedResourceGroupName {
		t.Fatalf("Expected the step to source 'constants.ArmResourceGroupName' from the state bag, but it did not.")
	}
}

func DeleteTestStateBagStepDeleteResourceGroup() multistep.StateBag {
	stateBag := new(multistep.BasicStateBag)
	stateBag.Put(constants.ArmResourceGroupName, "Unit Test: ResourceGroupName")

	return stateBag
}
