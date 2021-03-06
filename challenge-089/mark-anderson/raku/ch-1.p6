multi MAIN(UInt $n where * > 0) { #= greater than 0
    say gcd-sum($n);
}

multi MAIN {
    use Test;
    plan 2;

    ok gcd-sum(3) == 3, "Example 1";
    ok gcd-sum(4) == 7, "Example 2";
}

sub gcd-sum(UInt $n) {
    sum (1..$n).combinations(2).map(-> ($a, $b) { $a gcd $b });
}
