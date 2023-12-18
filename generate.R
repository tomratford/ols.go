N <- 40

set.seed(23101963)

x1 <- rnorm(N, 50, 5)
x2 <- rnorm(N, 4, 2)
x3 <- rnorm(N, -10, 8)
err <- rnorm(N, 0, 8)

betas <- c(0.45, 1.32, -3.78)

y <- betas[1] * x1 + betas[2] * x2 + betas[3] * x3 + err

data <- data.frame(y, x1, x2, x3)

# Check diagnostics
# lm(y ~ x1 + x2 + x3, data) |> summary()

write.csv(data, "input.csv")