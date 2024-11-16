# Greet
0. Create a repo with name `greet`
1. Create a directory
`mkdir greet`
2. Initilise the go package
`go mod init github.com/Crazy-Gopher/greet`
3. Write the code for your package
4. write the unittests and test it
`go test ./...` or `go test .`
5. Publish your package
```
git init
git add .
git commit -m "Initial commit"
git branch -M main
git remote add origin https://github.com/Crazy-Gopher/greet.git
git push origin main
```
6. Add tag to your code
```git tag v1.0.0```
```git push origin v1.0.0```

7. Go to client code and call function this package

8. Upgrade the version
```
git add .
git commit -m "Added HelloUpper"
git push origin main
git tag v1.0.2
git push origin v1.0.2
```