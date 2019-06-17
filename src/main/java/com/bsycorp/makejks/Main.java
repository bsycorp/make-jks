package com.bsycorp.makejks;

import picocli.CommandLine;

import java.io.File;
import java.io.StringReader;
import java.security.Security;
import java.util.Arrays;
import java.util.Collection;
import java.util.stream.Stream;
import org.apache.commons.io.FileUtils;
import org.bouncycastle.cert.X509CertificateHolder;
import org.bouncycastle.jce.provider.BouncyCastleProvider;
import org.bouncycastle.openssl.PEMParser;

public class Main {

    static {
        //Add BC to avoid using SunEC
        Security.addProvider(new BouncyCastleProvider());

        try {
            //Run keytool once to initialise the resource bundle, can remove if issue fixed
            sun.security.tools.keytool.Main.main(new String[]{"-help"});
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    public static void main(String[] args) throws Exception {
        ExtraArguments arguments = CommandLine.populateCommand(new ExtraArguments(), args);
        if(arguments.getDirectoryInput() != null){
            Collection<File> pemFiles = FileUtils.listFiles(arguments.getDirectoryInput(), new String[]{"pem"}, false);
            for (File pem : pemFiles){
                try (StringReader reader = new StringReader(FileUtils.readFileToString(pem, "UTF-8")); PEMParser pemParser = new PEMParser(reader)) {
                    X509CertificateHolder x509Cert = ((X509CertificateHolder) pemParser.readObject());
                    if(x509Cert == null) continue; //read object returns null for some parsing cases
                    
                    //create temp file for der
                    File tempDerFile = File.createTempFile("cert", ".der");
                    tempDerFile.deleteOnExit();

                    //write der encoded cert to temp file
                    FileUtils.writeByteArrayToFile(tempDerFile, x509Cert.getEncoded());

                    //call with single der file as arguments
                    String[] newArgs = Stream.concat(Arrays.stream(arguments.getRegularArguments()), Arrays.stream(new String[] {"-alias", pem.getName(), "-noprompt", "-file", tempDerFile.getAbsolutePath()}))
                            .toArray(String[]::new);
                    sun.security.tools.keytool.Main.main(newArgs);
                }
            }


        } else {
            sun.security.tools.keytool.Main.main(args);
        }
    }

}
