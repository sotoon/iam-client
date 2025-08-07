# Contributing to IAM Client

Thank you for your interest in contributing to the IAM Client! We welcome contributions from the community.

## Getting Started

1. Fork the repository
2. Clone your fork: `git clone https://github.com/YOUR_USERNAME/iam-client.git`
3. Create a new branch: `git checkout -b feature/your-feature-name`
4. Make your changes
5. Test your changes
6. Commit your changes: `git commit -m "Add your commit message"`
7. Push to your branch: `git push origin feature/your-feature-name`
8. Create a Pull Request

## Development Setup

### Prerequisites
- Go 1.19 or later
- Git

### Setup
```bash
# Clone the repository
git clone https://github.com/sotoon/iam-client.git
cd iam-client

# Install dependencies
go mod tidy
go mod vendor

# Run tests
go test ./...

# Run benchmarks
./benchmark.bash
```

## Code Guidelines

### Code Style
- Follow standard Go conventions and formatting
- Use `gofmt` to format your code
- Write clear, descriptive commit messages
- Add comments for complex logic

### Testing
- Write tests for new functionality
- Ensure all tests pass before submitting a PR
- Maintain or improve code coverage

### Documentation
- Update documentation for any API changes
- Include examples for new features
- Keep the README.md updated

## Pull Request Process

1. Ensure your code follows the project's coding standards
2. Update the README.md with details of changes if applicable
3. Add tests for new functionality
4. Ensure all tests pass
5. Update documentation as needed
6. Your PR will be reviewed by maintainers

## Reporting Issues

When reporting issues, please include:
- Go version
- Operating system
- Steps to reproduce the issue
- Expected behavior
- Actual behavior
- Any relevant logs or error messages

## Code of Conduct

Please note that this project is released with a Contributor Code of Conduct. By participating in this project you agree to abide by its terms.

## Questions?

If you have questions about contributing, please open an issue or reach out to the maintainers.

Thank you for contributing! 🎉
