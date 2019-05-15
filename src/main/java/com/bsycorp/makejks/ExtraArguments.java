package com.bsycorp.makejks;

import picocli.CommandLine;

import java.io.File;

public class ExtraArguments {

    @CommandLine.Option(names = {"-dir"}, required = false)
    private File directoryInput;

    @CommandLine.Option(names = {"-alias"}, required = false)
    private String alias;

    @CommandLine.Unmatched
    private String[] regularArguments;

    public File getDirectoryInput() {
        return directoryInput;
    }

    public String[] getRegularArguments() {
        return regularArguments;
    }
}
