# Security Policy

## Supported Versions

We release patches for security vulnerabilities. Which versions are eligible for receiving such patches depends on the CVSS v3.0 Rating:

| Version | Supported          |
| ------- | ------------------ |
| 1.0.x   | :white_check_mark: |
| < 1.0   | :x:                |

## Reporting a Vulnerability

The IAM Client team and community take security bugs seriously. We appreciate your efforts to responsibly disclose your findings, and will make every effort to acknowledge your contributions.

To report a security issue, please use the GitHub Security Advisory ["Report a Vulnerability"](https://github.com/sotoon/iam-client/security/advisories/new) tab.

The IAM Client team will send a response indicating the next steps in handling your report. After the initial reply to your report, the security team will keep you informed of the progress towards a fix and full announcement, and may ask for additional information or guidance.

## Security Best Practices

When using the IAM Client library:

1. **Never hardcode credentials** - Use environment variables or secure configuration management
2. **Use HTTPS** - Always use secure connections when communicating with IAM services
3. **Validate tokens** - Always validate access tokens before using them
4. **Keep dependencies updated** - Regularly update the library and its dependencies
5. **Follow principle of least privilege** - Only request the minimum permissions needed

## Security Features

The IAM Client includes several security features:

- Token validation and verification
- Secure HTTP client configuration
- Input sanitization and validation
- Error handling that doesn't leak sensitive information

## Disclosure Policy

When the security team receives a security bug report, they will assign it to a primary handler. This person will coordinate the fix and release process, involving the following steps:

- Confirm the problem and determine the affected versions
- Audit code to find any potential similar problems
- Prepare fixes for all releases still under maintenance
- Release new versions as soon as possible

## Comments on this Policy

If you have suggestions on how this process could be improved please submit a pull request.
