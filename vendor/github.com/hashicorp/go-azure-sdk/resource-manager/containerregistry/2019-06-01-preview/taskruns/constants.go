package taskruns

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Architecture string

const (
	ArchitectureAmdSixFour    Architecture = "amd64"
	ArchitectureArm           Architecture = "arm"
	ArchitectureArmSixFour    Architecture = "arm64"
	ArchitectureThreeEightSix Architecture = "386"
	ArchitectureXEightSix     Architecture = "x86"
)

func PossibleValuesForArchitecture() []string {
	return []string{
		string(ArchitectureAmdSixFour),
		string(ArchitectureArm),
		string(ArchitectureArmSixFour),
		string(ArchitectureThreeEightSix),
		string(ArchitectureXEightSix),
	}
}

func parseArchitecture(input string) (*Architecture, error) {
	vals := map[string]Architecture{
		"amd64": ArchitectureAmdSixFour,
		"arm":   ArchitectureArm,
		"arm64": ArchitectureArmSixFour,
		"386":   ArchitectureThreeEightSix,
		"x86":   ArchitectureXEightSix,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Architecture(input)
	return &out, nil
}

type OS string

const (
	OSLinux   OS = "Linux"
	OSWindows OS = "Windows"
)

func PossibleValuesForOS() []string {
	return []string{
		string(OSLinux),
		string(OSWindows),
	}
}

func parseOS(input string) (*OS, error) {
	vals := map[string]OS{
		"linux":   OSLinux,
		"windows": OSWindows,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OS(input)
	return &out, nil
}

type ProvisioningState string

const (
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCanceled),
		string(ProvisioningStateCreating),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateSucceeded),
		string(ProvisioningStateUpdating),
	}
}

func parseProvisioningState(input string) (*ProvisioningState, error) {
	vals := map[string]ProvisioningState{
		"canceled":  ProvisioningStateCanceled,
		"creating":  ProvisioningStateCreating,
		"deleting":  ProvisioningStateDeleting,
		"failed":    ProvisioningStateFailed,
		"succeeded": ProvisioningStateSucceeded,
		"updating":  ProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningState(input)
	return &out, nil
}

type RunStatus string

const (
	RunStatusCanceled  RunStatus = "Canceled"
	RunStatusError     RunStatus = "Error"
	RunStatusFailed    RunStatus = "Failed"
	RunStatusQueued    RunStatus = "Queued"
	RunStatusRunning   RunStatus = "Running"
	RunStatusStarted   RunStatus = "Started"
	RunStatusSucceeded RunStatus = "Succeeded"
	RunStatusTimeout   RunStatus = "Timeout"
)

func PossibleValuesForRunStatus() []string {
	return []string{
		string(RunStatusCanceled),
		string(RunStatusError),
		string(RunStatusFailed),
		string(RunStatusQueued),
		string(RunStatusRunning),
		string(RunStatusStarted),
		string(RunStatusSucceeded),
		string(RunStatusTimeout),
	}
}

func parseRunStatus(input string) (*RunStatus, error) {
	vals := map[string]RunStatus{
		"canceled":  RunStatusCanceled,
		"error":     RunStatusError,
		"failed":    RunStatusFailed,
		"queued":    RunStatusQueued,
		"running":   RunStatusRunning,
		"started":   RunStatusStarted,
		"succeeded": RunStatusSucceeded,
		"timeout":   RunStatusTimeout,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RunStatus(input)
	return &out, nil
}

type RunType string

const (
	RunTypeAutoBuild  RunType = "AutoBuild"
	RunTypeAutoRun    RunType = "AutoRun"
	RunTypeQuickBuild RunType = "QuickBuild"
	RunTypeQuickRun   RunType = "QuickRun"
)

func PossibleValuesForRunType() []string {
	return []string{
		string(RunTypeAutoBuild),
		string(RunTypeAutoRun),
		string(RunTypeQuickBuild),
		string(RunTypeQuickRun),
	}
}

func parseRunType(input string) (*RunType, error) {
	vals := map[string]RunType{
		"autobuild":  RunTypeAutoBuild,
		"autorun":    RunTypeAutoRun,
		"quickbuild": RunTypeQuickBuild,
		"quickrun":   RunTypeQuickRun,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RunType(input)
	return &out, nil
}

type SecretObjectType string

const (
	SecretObjectTypeOpaque      SecretObjectType = "Opaque"
	SecretObjectTypeVaultsecret SecretObjectType = "Vaultsecret"
)

func PossibleValuesForSecretObjectType() []string {
	return []string{
		string(SecretObjectTypeOpaque),
		string(SecretObjectTypeVaultsecret),
	}
}

func parseSecretObjectType(input string) (*SecretObjectType, error) {
	vals := map[string]SecretObjectType{
		"opaque":      SecretObjectTypeOpaque,
		"vaultsecret": SecretObjectTypeVaultsecret,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecretObjectType(input)
	return &out, nil
}

type SourceRegistryLoginMode string

const (
	SourceRegistryLoginModeDefault SourceRegistryLoginMode = "Default"
	SourceRegistryLoginModeNone    SourceRegistryLoginMode = "None"
)

func PossibleValuesForSourceRegistryLoginMode() []string {
	return []string{
		string(SourceRegistryLoginModeDefault),
		string(SourceRegistryLoginModeNone),
	}
}

func parseSourceRegistryLoginMode(input string) (*SourceRegistryLoginMode, error) {
	vals := map[string]SourceRegistryLoginMode{
		"default": SourceRegistryLoginModeDefault,
		"none":    SourceRegistryLoginModeNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SourceRegistryLoginMode(input)
	return &out, nil
}

type Variant string

const (
	VariantVEight Variant = "v8"
	VariantVSeven Variant = "v7"
	VariantVSix   Variant = "v6"
)

func PossibleValuesForVariant() []string {
	return []string{
		string(VariantVEight),
		string(VariantVSeven),
		string(VariantVSix),
	}
}

func parseVariant(input string) (*Variant, error) {
	vals := map[string]Variant{
		"v8": VariantVEight,
		"v7": VariantVSeven,
		"v6": VariantVSix,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Variant(input)
	return &out, nil
}
