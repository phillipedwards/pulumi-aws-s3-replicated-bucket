// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export namespace crd.k8s.amazonaws.com {
    export namespace v1alpha1 {
        export interface ENIConfigSpec {
            securityGroups?: string[];
            subnet?: string;
        }

    }
}
