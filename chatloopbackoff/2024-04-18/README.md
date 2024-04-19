## Episode 4: OpenFeature

YouTube: ####
### News

### Summary

- What is OpenFeature?
    - What is a feature flag?
      Feature flags are a software development technique that allows teams to enable, disable or change the behavior of certain features or code paths in a product or service, without modifying the source code.
    - What is OpenFeature?
      OpenFeature is an open specification that provides a vendor-agnostic, community-driven API for feature flagging that works with your favorite feature flag management tool or in-house solution.
    - Why OpenFeature?
      It is a vendor agnostic way to define feature flags, and query them for use in your application by following a pre-defined standard
- 

- [Getting started with OpenFeature](https://openfeature.dev/docs/tutorials/getting-started/go)
  - Created basic golang app
  - flagd spun up via podman (Because Jeefy's docker instance was sad.)
  - Connected the two with a minor code change
  - Defined feature flags we want served
- OpenFeature native Kubernetes Support
  - CRD that defines Feature Flags and serves to internal apps

### TODO / Questions / Musings
- [PR to update the OpenFeature docs](https://github.com/open-feature/openfeature.dev/pull/525)
  - Example Pt 2 in OpenFeature docs can't be fully copied, missing end curly brace for main
- Revisit more [cloud native example](https://openfeature.dev/docs/tutorials/ofo) with Feature Flags as CRDs
