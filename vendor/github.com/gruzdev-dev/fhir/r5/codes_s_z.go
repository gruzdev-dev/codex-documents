package models

// SPDXLicense represents codes from http://hl7.org/fhir/ValueSet/spdx-license
type SPDXLicense string

const (
	SPDXLicenseNotOpenSource                  SPDXLicense = "not-open-source"
	SPDXLicenseCode0BSD                       SPDXLicense = "0BSD"
	SPDXLicenseAAL                            SPDXLicense = "AAL"
	SPDXLicenseAbstyles                       SPDXLicense = "Abstyles"
	SPDXLicenseAdobe2006                      SPDXLicense = "Adobe-2006"
	SPDXLicenseAdobeGlyph                     SPDXLicense = "Adobe-Glyph"
	SPDXLicenseADSL                           SPDXLicense = "ADSL"
	SPDXLicenseAFL11                          SPDXLicense = "AFL-1.1"
	SPDXLicenseAFL12                          SPDXLicense = "AFL-1.2"
	SPDXLicenseAFL20                          SPDXLicense = "AFL-2.0"
	SPDXLicenseAFL21                          SPDXLicense = "AFL-2.1"
	SPDXLicenseAFL30                          SPDXLicense = "AFL-3.0"
	SPDXLicenseAfmparse                       SPDXLicense = "Afmparse"
	SPDXLicenseAGPL10Only                     SPDXLicense = "AGPL-1.0-only"
	SPDXLicenseAGPL10OrLater                  SPDXLicense = "AGPL-1.0-or-later"
	SPDXLicenseAGPL30Only                     SPDXLicense = "AGPL-3.0-only"
	SPDXLicenseAGPL30OrLater                  SPDXLicense = "AGPL-3.0-or-later"
	SPDXLicenseAladdin                        SPDXLicense = "Aladdin"
	SPDXLicenseAMDPLPA                        SPDXLicense = "AMDPLPA"
	SPDXLicenseAML                            SPDXLicense = "AML"
	SPDXLicenseAMPAS                          SPDXLicense = "AMPAS"
	SPDXLicenseANTLRPD                        SPDXLicense = "ANTLR-PD"
	SPDXLicenseApache10                       SPDXLicense = "Apache-1.0"
	SPDXLicenseApache11                       SPDXLicense = "Apache-1.1"
	SPDXLicenseApache20                       SPDXLicense = "Apache-2.0"
	SPDXLicenseAPAFML                         SPDXLicense = "APAFML"
	SPDXLicenseAPL10                          SPDXLicense = "APL-1.0"
	SPDXLicenseAPSL10                         SPDXLicense = "APSL-1.0"
	SPDXLicenseAPSL11                         SPDXLicense = "APSL-1.1"
	SPDXLicenseAPSL12                         SPDXLicense = "APSL-1.2"
	SPDXLicenseAPSL20                         SPDXLicense = "APSL-2.0"
	SPDXLicenseArtistic10Cl8                  SPDXLicense = "Artistic-1.0-cl8"
	SPDXLicenseArtistic10Perl                 SPDXLicense = "Artistic-1.0-Perl"
	SPDXLicenseArtistic10                     SPDXLicense = "Artistic-1.0"
	SPDXLicenseArtistic20                     SPDXLicense = "Artistic-2.0"
	SPDXLicenseBahyph                         SPDXLicense = "Bahyph"
	SPDXLicenseBarr                           SPDXLicense = "Barr"
	SPDXLicenseBeerware                       SPDXLicense = "Beerware"
	SPDXLicenseBitTorrent10                   SPDXLicense = "BitTorrent-1.0"
	SPDXLicenseBitTorrent11                   SPDXLicense = "BitTorrent-1.1"
	SPDXLicenseBorceux                        SPDXLicense = "Borceux"
	SPDXLicenseBSD1Clause                     SPDXLicense = "BSD-1-Clause"
	SPDXLicenseBSD2ClauseFreeBSD              SPDXLicense = "BSD-2-Clause-FreeBSD"
	SPDXLicenseBSD2ClauseNetBSD               SPDXLicense = "BSD-2-Clause-NetBSD"
	SPDXLicenseBSD2ClausePatent               SPDXLicense = "BSD-2-Clause-Patent"
	SPDXLicenseBSD2Clause                     SPDXLicense = "BSD-2-Clause"
	SPDXLicenseBSD3ClauseAttribution          SPDXLicense = "BSD-3-Clause-Attribution"
	SPDXLicenseBSD3ClauseClear                SPDXLicense = "BSD-3-Clause-Clear"
	SPDXLicenseBSD3ClauseLBNL                 SPDXLicense = "BSD-3-Clause-LBNL"
	SPDXLicenseBSD3ClauseNoNuclearLicense2014 SPDXLicense = "BSD-3-Clause-No-Nuclear-License-2014"
	SPDXLicenseBSD3ClauseNoNuclearLicense     SPDXLicense = "BSD-3-Clause-No-Nuclear-License"
	SPDXLicenseBSD3ClauseNoNuclearWarranty    SPDXLicense = "BSD-3-Clause-No-Nuclear-Warranty"
	SPDXLicenseBSD3Clause                     SPDXLicense = "BSD-3-Clause"
	SPDXLicenseBSD4ClauseUC                   SPDXLicense = "BSD-4-Clause-UC"
	SPDXLicenseBSD4Clause                     SPDXLicense = "BSD-4-Clause"
	SPDXLicenseBSDProtection                  SPDXLicense = "BSD-Protection"
	SPDXLicenseBSDSourceCode                  SPDXLicense = "BSD-Source-Code"
	SPDXLicenseBSL10                          SPDXLicense = "BSL-1.0"
	SPDXLicenseBzip2105                       SPDXLicense = "bzip2-1.0.5"
	SPDXLicenseBzip2106                       SPDXLicense = "bzip2-1.0.6"
	SPDXLicenseCaldera                        SPDXLicense = "Caldera"
	SPDXLicenseCATOSL11                       SPDXLicense = "CATOSL-1.1"
	SPDXLicenseCCBY10                         SPDXLicense = "CC-BY-1.0"
	SPDXLicenseCCBY20                         SPDXLicense = "CC-BY-2.0"
	SPDXLicenseCCBY25                         SPDXLicense = "CC-BY-2.5"
	SPDXLicenseCCBY30                         SPDXLicense = "CC-BY-3.0"
	SPDXLicenseCCBY40                         SPDXLicense = "CC-BY-4.0"
	SPDXLicenseCCBYNC10                       SPDXLicense = "CC-BY-NC-1.0"
	SPDXLicenseCCBYNC20                       SPDXLicense = "CC-BY-NC-2.0"
	SPDXLicenseCCBYNC25                       SPDXLicense = "CC-BY-NC-2.5"
	SPDXLicenseCCBYNC30                       SPDXLicense = "CC-BY-NC-3.0"
	SPDXLicenseCCBYNC40                       SPDXLicense = "CC-BY-NC-4.0"
	SPDXLicenseCCBYNCND10                     SPDXLicense = "CC-BY-NC-ND-1.0"
	SPDXLicenseCCBYNCND20                     SPDXLicense = "CC-BY-NC-ND-2.0"
	SPDXLicenseCCBYNCND25                     SPDXLicense = "CC-BY-NC-ND-2.5"
	SPDXLicenseCCBYNCND30                     SPDXLicense = "CC-BY-NC-ND-3.0"
	SPDXLicenseCCBYNCND40                     SPDXLicense = "CC-BY-NC-ND-4.0"
	SPDXLicenseCCBYNCSA10                     SPDXLicense = "CC-BY-NC-SA-1.0"
	SPDXLicenseCCBYNCSA20                     SPDXLicense = "CC-BY-NC-SA-2.0"
	SPDXLicenseCCBYNCSA25                     SPDXLicense = "CC-BY-NC-SA-2.5"
	SPDXLicenseCCBYNCSA30                     SPDXLicense = "CC-BY-NC-SA-3.0"
	SPDXLicenseCCBYNCSA40                     SPDXLicense = "CC-BY-NC-SA-4.0"
	SPDXLicenseCCBYND10                       SPDXLicense = "CC-BY-ND-1.0"
	SPDXLicenseCCBYND20                       SPDXLicense = "CC-BY-ND-2.0"
	SPDXLicenseCCBYND25                       SPDXLicense = "CC-BY-ND-2.5"
	SPDXLicenseCCBYND30                       SPDXLicense = "CC-BY-ND-3.0"
	SPDXLicenseCCBYND40                       SPDXLicense = "CC-BY-ND-4.0"
	SPDXLicenseCCBYSA10                       SPDXLicense = "CC-BY-SA-1.0"
	SPDXLicenseCCBYSA20                       SPDXLicense = "CC-BY-SA-2.0"
	SPDXLicenseCCBYSA25                       SPDXLicense = "CC-BY-SA-2.5"
	SPDXLicenseCCBYSA30                       SPDXLicense = "CC-BY-SA-3.0"
	SPDXLicenseCCBYSA40                       SPDXLicense = "CC-BY-SA-4.0"
	SPDXLicenseCC010                          SPDXLicense = "CC0-1.0"
	SPDXLicenseCDDL10                         SPDXLicense = "CDDL-1.0"
	SPDXLicenseCDDL11                         SPDXLicense = "CDDL-1.1"
	SPDXLicenseCDLAPermissive10               SPDXLicense = "CDLA-Permissive-1.0"
	SPDXLicenseCDLASharing10                  SPDXLicense = "CDLA-Sharing-1.0"
	SPDXLicenseCECILL10                       SPDXLicense = "CECILL-1.0"
	SPDXLicenseCECILL11                       SPDXLicense = "CECILL-1.1"
	SPDXLicenseCECILL20                       SPDXLicense = "CECILL-2.0"
	SPDXLicenseCECILL21                       SPDXLicense = "CECILL-2.1"
	SPDXLicenseCECILLB                        SPDXLicense = "CECILL-B"
	SPDXLicenseCECILLC                        SPDXLicense = "CECILL-C"
	SPDXLicenseClArtistic                     SPDXLicense = "ClArtistic"
	SPDXLicenseCNRIJython                     SPDXLicense = "CNRI-Jython"
	SPDXLicenseCNRIPythonGPLCompatible        SPDXLicense = "CNRI-Python-GPL-Compatible"
	SPDXLicenseCNRIPython                     SPDXLicense = "CNRI-Python"
	SPDXLicenseCondor11                       SPDXLicense = "Condor-1.1"
	SPDXLicenseCPAL10                         SPDXLicense = "CPAL-1.0"
	SPDXLicenseCPL10                          SPDXLicense = "CPL-1.0"
	SPDXLicenseCPOL102                        SPDXLicense = "CPOL-1.02"
	SPDXLicenseCrossword                      SPDXLicense = "Crossword"
	SPDXLicenseCrystalStacker                 SPDXLicense = "CrystalStacker"
	SPDXLicenseCUAOPL10                       SPDXLicense = "CUA-OPL-1.0"
	SPDXLicenseCube                           SPDXLicense = "Cube"
	SPDXLicenseCurl                           SPDXLicense = "curl"
	SPDXLicenseDFSL10                         SPDXLicense = "D-FSL-1.0"
	SPDXLicenseDiffmark                       SPDXLicense = "diffmark"
	SPDXLicenseDOC                            SPDXLicense = "DOC"
	SPDXLicenseDotseqn                        SPDXLicense = "Dotseqn"
	SPDXLicenseDSDP                           SPDXLicense = "DSDP"
	SPDXLicenseDvipdfm                        SPDXLicense = "dvipdfm"
	SPDXLicenseECL10                          SPDXLicense = "ECL-1.0"
	SPDXLicenseECL20                          SPDXLicense = "ECL-2.0"
	SPDXLicenseEFL10                          SPDXLicense = "EFL-1.0"
	SPDXLicenseEFL20                          SPDXLicense = "EFL-2.0"
	SPDXLicenseEGenix                         SPDXLicense = "eGenix"
	SPDXLicenseEntessa                        SPDXLicense = "Entessa"
	SPDXLicenseEPL10                          SPDXLicense = "EPL-1.0"
	SPDXLicenseEPL20                          SPDXLicense = "EPL-2.0"
	SPDXLicenseErlPL11                        SPDXLicense = "ErlPL-1.1"
	SPDXLicenseEUDatagrid                     SPDXLicense = "EUDatagrid"
	SPDXLicenseEUPL10                         SPDXLicense = "EUPL-1.0"
	SPDXLicenseEUPL11                         SPDXLicense = "EUPL-1.1"
	SPDXLicenseEUPL12                         SPDXLicense = "EUPL-1.2"
	SPDXLicenseEurosym                        SPDXLicense = "Eurosym"
	SPDXLicenseFair                           SPDXLicense = "Fair"
	SPDXLicenseFrameworx10                    SPDXLicense = "Frameworx-1.0"
	SPDXLicenseFreeImage                      SPDXLicense = "FreeImage"
	SPDXLicenseFSFAP                          SPDXLicense = "FSFAP"
	SPDXLicenseFSFUL                          SPDXLicense = "FSFUL"
	SPDXLicenseFSFULLR                        SPDXLicense = "FSFULLR"
	SPDXLicenseFTL                            SPDXLicense = "FTL"
	SPDXLicenseGFDL11Only                     SPDXLicense = "GFDL-1.1-only"
	SPDXLicenseGFDL11OrLater                  SPDXLicense = "GFDL-1.1-or-later"
	SPDXLicenseGFDL12Only                     SPDXLicense = "GFDL-1.2-only"
	SPDXLicenseGFDL12OrLater                  SPDXLicense = "GFDL-1.2-or-later"
	SPDXLicenseGFDL13Only                     SPDXLicense = "GFDL-1.3-only"
	SPDXLicenseGFDL13OrLater                  SPDXLicense = "GFDL-1.3-or-later"
	SPDXLicenseGiftware                       SPDXLicense = "Giftware"
	SPDXLicenseGL2PS                          SPDXLicense = "GL2PS"
	SPDXLicenseGlide                          SPDXLicense = "Glide"
	SPDXLicenseGlulxe                         SPDXLicense = "Glulxe"
	SPDXLicenseGnuplot                        SPDXLicense = "gnuplot"
	SPDXLicenseGPL10Only                      SPDXLicense = "GPL-1.0-only"
	SPDXLicenseGPL10OrLater                   SPDXLicense = "GPL-1.0-or-later"
	SPDXLicenseGPL20Only                      SPDXLicense = "GPL-2.0-only"
	SPDXLicenseGPL20OrLater                   SPDXLicense = "GPL-2.0-or-later"
	SPDXLicenseGPL30Only                      SPDXLicense = "GPL-3.0-only"
	SPDXLicenseGPL30OrLater                   SPDXLicense = "GPL-3.0-or-later"
	SPDXLicenseGSOAP13b                       SPDXLicense = "gSOAP-1.3b"
	SPDXLicenseHaskellReport                  SPDXLicense = "HaskellReport"
	SPDXLicenseHPND                           SPDXLicense = "HPND"
	SPDXLicenseIBMPibs                        SPDXLicense = "IBM-pibs"
	SPDXLicenseICU                            SPDXLicense = "ICU"
	SPDXLicenseIJG                            SPDXLicense = "IJG"
	SPDXLicenseImageMagick                    SPDXLicense = "ImageMagick"
	SPDXLicenseIMatix                         SPDXLicense = "iMatix"
	SPDXLicenseImlib2                         SPDXLicense = "Imlib2"
	SPDXLicenseInfoZIP                        SPDXLicense = "Info-ZIP"
	SPDXLicenseIntelACPI                      SPDXLicense = "Intel-ACPI"
	SPDXLicenseIntel                          SPDXLicense = "Intel"
	SPDXLicenseInterbase10                    SPDXLicense = "Interbase-1.0"
	SPDXLicenseIPA                            SPDXLicense = "IPA"
	SPDXLicenseIPL10                          SPDXLicense = "IPL-1.0"
	SPDXLicenseISC                            SPDXLicense = "ISC"
	SPDXLicenseJasPer20                       SPDXLicense = "JasPer-2.0"
	SPDXLicenseJSON                           SPDXLicense = "JSON"
	SPDXLicenseLAL12                          SPDXLicense = "LAL-1.2"
	SPDXLicenseLAL13                          SPDXLicense = "LAL-1.3"
	SPDXLicenseLatex2e                        SPDXLicense = "Latex2e"
	SPDXLicenseLeptonica                      SPDXLicense = "Leptonica"
	SPDXLicenseLGPL20Only                     SPDXLicense = "LGPL-2.0-only"
	SPDXLicenseLGPL20OrLater                  SPDXLicense = "LGPL-2.0-or-later"
	SPDXLicenseLGPL21Only                     SPDXLicense = "LGPL-2.1-only"
	SPDXLicenseLGPL21OrLater                  SPDXLicense = "LGPL-2.1-or-later"
	SPDXLicenseLGPL30Only                     SPDXLicense = "LGPL-3.0-only"
	SPDXLicenseLGPL30OrLater                  SPDXLicense = "LGPL-3.0-or-later"
	SPDXLicenseLGPLLR                         SPDXLicense = "LGPLLR"
	SPDXLicenseLibpng                         SPDXLicense = "Libpng"
	SPDXLicenseLibtiff                        SPDXLicense = "libtiff"
	SPDXLicenseLiLiQP11                       SPDXLicense = "LiLiQ-P-1.1"
	SPDXLicenseLiLiQR11                       SPDXLicense = "LiLiQ-R-1.1"
	SPDXLicenseLiLiQRplus11                   SPDXLicense = "LiLiQ-Rplus-1.1"
	SPDXLicenseLinuxOpenIB                    SPDXLicense = "Linux-OpenIB"
	SPDXLicenseLPL10                          SPDXLicense = "LPL-1.0"
	SPDXLicenseLPL102                         SPDXLicense = "LPL-1.02"
	SPDXLicenseLPPL10                         SPDXLicense = "LPPL-1.0"
	SPDXLicenseLPPL11                         SPDXLicense = "LPPL-1.1"
	SPDXLicenseLPPL12                         SPDXLicense = "LPPL-1.2"
	SPDXLicenseLPPL13a                        SPDXLicense = "LPPL-1.3a"
	SPDXLicenseLPPL13c                        SPDXLicense = "LPPL-1.3c"
	SPDXLicenseMakeIndex                      SPDXLicense = "MakeIndex"
	SPDXLicenseMirOS                          SPDXLicense = "MirOS"
	SPDXLicenseMIT0                           SPDXLicense = "MIT-0"
	SPDXLicenseMITAdvertising                 SPDXLicense = "MIT-advertising"
	SPDXLicenseMITCMU                         SPDXLicense = "MIT-CMU"
	SPDXLicenseMITEnna                        SPDXLicense = "MIT-enna"
	SPDXLicenseMITFeh                         SPDXLicense = "MIT-feh"
	SPDXLicenseMIT                            SPDXLicense = "MIT"
	SPDXLicenseMITNFA                         SPDXLicense = "MITNFA"
	SPDXLicenseMotosoto                       SPDXLicense = "Motosoto"
	SPDXLicenseMpich2                         SPDXLicense = "mpich2"
	SPDXLicenseMPL10                          SPDXLicense = "MPL-1.0"
	SPDXLicenseMPL11                          SPDXLicense = "MPL-1.1"
	SPDXLicenseMPL20NoCopyleftException       SPDXLicense = "MPL-2.0-no-copyleft-exception"
	SPDXLicenseMPL20                          SPDXLicense = "MPL-2.0"
	SPDXLicenseMSPL                           SPDXLicense = "MS-PL"
	SPDXLicenseMSRL                           SPDXLicense = "MS-RL"
	SPDXLicenseMTLL                           SPDXLicense = "MTLL"
	SPDXLicenseMultics                        SPDXLicense = "Multics"
	SPDXLicenseMup                            SPDXLicense = "Mup"
	SPDXLicenseNASA13                         SPDXLicense = "NASA-1.3"
	SPDXLicenseNaumen                         SPDXLicense = "Naumen"
	SPDXLicenseNBPL10                         SPDXLicense = "NBPL-1.0"
	SPDXLicenseNCSA                           SPDXLicense = "NCSA"
	SPDXLicenseNetSNMP                        SPDXLicense = "Net-SNMP"
	SPDXLicenseNetCDF                         SPDXLicense = "NetCDF"
	SPDXLicenseNewsletr                       SPDXLicense = "Newsletr"
	SPDXLicenseNGPL                           SPDXLicense = "NGPL"
	SPDXLicenseNLOD10                         SPDXLicense = "NLOD-1.0"
	SPDXLicenseNLPL                           SPDXLicense = "NLPL"
	SPDXLicenseNokia                          SPDXLicense = "Nokia"
	SPDXLicenseNOSL                           SPDXLicense = "NOSL"
	SPDXLicenseNoweb                          SPDXLicense = "Noweb"
	SPDXLicenseNPL10                          SPDXLicense = "NPL-1.0"
	SPDXLicenseNPL11                          SPDXLicense = "NPL-1.1"
	SPDXLicenseNPOSL30                        SPDXLicense = "NPOSL-3.0"
	SPDXLicenseNRL                            SPDXLicense = "NRL"
	SPDXLicenseNTP                            SPDXLicense = "NTP"
	SPDXLicenseOCCTPL                         SPDXLicense = "OCCT-PL"
	SPDXLicenseOCLC20                         SPDXLicense = "OCLC-2.0"
	SPDXLicenseODbL10                         SPDXLicense = "ODbL-1.0"
	SPDXLicenseOFL10                          SPDXLicense = "OFL-1.0"
	SPDXLicenseOFL11                          SPDXLicense = "OFL-1.1"
	SPDXLicenseOGTSL                          SPDXLicense = "OGTSL"
	SPDXLicenseOLDAP11                        SPDXLicense = "OLDAP-1.1"
	SPDXLicenseOLDAP12                        SPDXLicense = "OLDAP-1.2"
	SPDXLicenseOLDAP13                        SPDXLicense = "OLDAP-1.3"
	SPDXLicenseOLDAP14                        SPDXLicense = "OLDAP-1.4"
	SPDXLicenseOLDAP201                       SPDXLicense = "OLDAP-2.0.1"
	SPDXLicenseOLDAP20                        SPDXLicense = "OLDAP-2.0"
	SPDXLicenseOLDAP21                        SPDXLicense = "OLDAP-2.1"
	SPDXLicenseOLDAP221                       SPDXLicense = "OLDAP-2.2.1"
	SPDXLicenseOLDAP222                       SPDXLicense = "OLDAP-2.2.2"
	SPDXLicenseOLDAP22                        SPDXLicense = "OLDAP-2.2"
	SPDXLicenseOLDAP23                        SPDXLicense = "OLDAP-2.3"
	SPDXLicenseOLDAP24                        SPDXLicense = "OLDAP-2.4"
	SPDXLicenseOLDAP25                        SPDXLicense = "OLDAP-2.5"
	SPDXLicenseOLDAP26                        SPDXLicense = "OLDAP-2.6"
	SPDXLicenseOLDAP27                        SPDXLicense = "OLDAP-2.7"
	SPDXLicenseOLDAP28                        SPDXLicense = "OLDAP-2.8"
	SPDXLicenseOML                            SPDXLicense = "OML"
	SPDXLicenseOpenSSL                        SPDXLicense = "OpenSSL"
	SPDXLicenseOPL10                          SPDXLicense = "OPL-1.0"
	SPDXLicenseOSETPL21                       SPDXLicense = "OSET-PL-2.1"
	SPDXLicenseOSL10                          SPDXLicense = "OSL-1.0"
	SPDXLicenseOSL11                          SPDXLicense = "OSL-1.1"
	SPDXLicenseOSL20                          SPDXLicense = "OSL-2.0"
	SPDXLicenseOSL21                          SPDXLicense = "OSL-2.1"
	SPDXLicenseOSL30                          SPDXLicense = "OSL-3.0"
	SPDXLicensePDDL10                         SPDXLicense = "PDDL-1.0"
	SPDXLicensePHP30                          SPDXLicense = "PHP-3.0"
	SPDXLicensePHP301                         SPDXLicense = "PHP-3.01"
	SPDXLicensePlexus                         SPDXLicense = "Plexus"
	SPDXLicensePostgreSQL                     SPDXLicense = "PostgreSQL"
	SPDXLicensePsfrag                         SPDXLicense = "psfrag"
	SPDXLicensePsutils                        SPDXLicense = "psutils"
	SPDXLicensePython20                       SPDXLicense = "Python-2.0"
	SPDXLicenseQhull                          SPDXLicense = "Qhull"
	SPDXLicenseQPL10                          SPDXLicense = "QPL-1.0"
	SPDXLicenseRdisc                          SPDXLicense = "Rdisc"
	SPDXLicenseRHeCos11                       SPDXLicense = "RHeCos-1.1"
	SPDXLicenseRPL11                          SPDXLicense = "RPL-1.1"
	SPDXLicenseRPL15                          SPDXLicense = "RPL-1.5"
	SPDXLicenseRPSL10                         SPDXLicense = "RPSL-1.0"
	SPDXLicenseRSAMD                          SPDXLicense = "RSA-MD"
	SPDXLicenseRSCPL                          SPDXLicense = "RSCPL"
	SPDXLicenseRuby                           SPDXLicense = "Ruby"
	SPDXLicenseSAXPD                          SPDXLicense = "SAX-PD"
	SPDXLicenseSaxpath                        SPDXLicense = "Saxpath"
	SPDXLicenseSCEA                           SPDXLicense = "SCEA"
	SPDXLicenseSendmail                       SPDXLicense = "Sendmail"
	SPDXLicenseSGIB10                         SPDXLicense = "SGI-B-1.0"
	SPDXLicenseSGIB11                         SPDXLicense = "SGI-B-1.1"
	SPDXLicenseSGIB20                         SPDXLicense = "SGI-B-2.0"
	SPDXLicenseSimPL20                        SPDXLicense = "SimPL-2.0"
	SPDXLicenseSISSL12                        SPDXLicense = "SISSL-1.2"
	SPDXLicenseSISSL                          SPDXLicense = "SISSL"
	SPDXLicenseSleepycat                      SPDXLicense = "Sleepycat"
	SPDXLicenseSMLNJ                          SPDXLicense = "SMLNJ"
	SPDXLicenseSMPPL                          SPDXLicense = "SMPPL"
	SPDXLicenseSNIA                           SPDXLicense = "SNIA"
	SPDXLicenseSpencer86                      SPDXLicense = "Spencer-86"
	SPDXLicenseSpencer94                      SPDXLicense = "Spencer-94"
	SPDXLicenseSpencer99                      SPDXLicense = "Spencer-99"
	SPDXLicenseSPL10                          SPDXLicense = "SPL-1.0"
	SPDXLicenseSugarCRM113                    SPDXLicense = "SugarCRM-1.1.3"
	SPDXLicenseSWL                            SPDXLicense = "SWL"
	SPDXLicenseTCL                            SPDXLicense = "TCL"
	SPDXLicenseTCPWrappers                    SPDXLicense = "TCP-wrappers"
	SPDXLicenseTMate                          SPDXLicense = "TMate"
	SPDXLicenseTORQUE11                       SPDXLicense = "TORQUE-1.1"
	SPDXLicenseTOSL                           SPDXLicense = "TOSL"
	SPDXLicenseUnicodeDFS2015                 SPDXLicense = "Unicode-DFS-2015"
	SPDXLicenseUnicodeDFS2016                 SPDXLicense = "Unicode-DFS-2016"
	SPDXLicenseUnicodeTOU                     SPDXLicense = "Unicode-TOU"
	SPDXLicenseUnlicense                      SPDXLicense = "Unlicense"
	SPDXLicenseUPL10                          SPDXLicense = "UPL-1.0"
	SPDXLicenseVim                            SPDXLicense = "Vim"
	SPDXLicenseVOSTROM                        SPDXLicense = "VOSTROM"
	SPDXLicenseVSL10                          SPDXLicense = "VSL-1.0"
	SPDXLicenseW3C19980720                    SPDXLicense = "W3C-19980720"
	SPDXLicenseW3C20150513                    SPDXLicense = "W3C-20150513"
	SPDXLicenseW3C                            SPDXLicense = "W3C"
	SPDXLicenseWatcom10                       SPDXLicense = "Watcom-1.0"
	SPDXLicenseWsuipa                         SPDXLicense = "Wsuipa"
	SPDXLicenseWTFPL                          SPDXLicense = "WTFPL"
	SPDXLicenseX11                            SPDXLicense = "X11"
	SPDXLicenseXerox                          SPDXLicense = "Xerox"
	SPDXLicenseXFree8611                      SPDXLicense = "XFree86-1.1"
	SPDXLicenseXinetd                         SPDXLicense = "xinetd"
	SPDXLicenseXnet                           SPDXLicense = "Xnet"
	SPDXLicenseXpp                            SPDXLicense = "xpp"
	SPDXLicenseXSkat                          SPDXLicense = "XSkat"
	SPDXLicenseYPL10                          SPDXLicense = "YPL-1.0"
	SPDXLicenseYPL11                          SPDXLicense = "YPL-1.1"
	SPDXLicenseZed                            SPDXLicense = "Zed"
	SPDXLicenseZend20                         SPDXLicense = "Zend-2.0"
	SPDXLicenseZimbra13                       SPDXLicense = "Zimbra-1.3"
	SPDXLicenseZimbra14                       SPDXLicense = "Zimbra-1.4"
	SPDXLicenseZlibAcknowledgement            SPDXLicense = "zlib-acknowledgement"
	SPDXLicenseZlib                           SPDXLicense = "Zlib"
	SPDXLicenseZPL11                          SPDXLicense = "ZPL-1.1"
	SPDXLicenseZPL20                          SPDXLicense = "ZPL-2.0"
	SPDXLicenseZPL21                          SPDXLicense = "ZPL-2.1"
)

// SearchComparator represents codes from http://hl7.org/fhir/ValueSet/search-comparator
type SearchComparator string

const (
	SearchComparatorEq SearchComparator = "eq"
	SearchComparatorNe SearchComparator = "ne"
	SearchComparatorGt SearchComparator = "gt"
	SearchComparatorLt SearchComparator = "lt"
	SearchComparatorGe SearchComparator = "ge"
	SearchComparatorLe SearchComparator = "le"
	SearchComparatorSa SearchComparator = "sa"
	SearchComparatorEb SearchComparator = "eb"
	SearchComparatorAp SearchComparator = "ap"
)

// SearchEntryMode represents codes from http://hl7.org/fhir/ValueSet/search-entry-mode
type SearchEntryMode string

const (
	SearchEntryModeMatch   SearchEntryMode = "match"
	SearchEntryModeInclude SearchEntryMode = "include"
	SearchEntryModeOutcome SearchEntryMode = "outcome"
)

// SearchModifierAllCodes represents codes from http://hl7.org/fhir/ValueSet/search-modifier-all-codes
type SearchModifierAllCodes string

const (
	SearchModifierAllCodesMissing      SearchModifierAllCodes = "missing"
	SearchModifierAllCodesExact        SearchModifierAllCodes = "exact"
	SearchModifierAllCodesContains     SearchModifierAllCodes = "contains"
	SearchModifierAllCodesNot          SearchModifierAllCodes = "not"
	SearchModifierAllCodesText         SearchModifierAllCodes = "text"
	SearchModifierAllCodesIn           SearchModifierAllCodes = "in"
	SearchModifierAllCodesNotIn        SearchModifierAllCodes = "not-in"
	SearchModifierAllCodesBelow        SearchModifierAllCodes = "below"
	SearchModifierAllCodesAbove        SearchModifierAllCodes = "above"
	SearchModifierAllCodesType         SearchModifierAllCodes = "type"
	SearchModifierAllCodesIdentifier   SearchModifierAllCodes = "identifier"
	SearchModifierAllCodesOfType       SearchModifierAllCodes = "of-type"
	SearchModifierAllCodesCodeText     SearchModifierAllCodes = "code-text"
	SearchModifierAllCodesTextAdvanced SearchModifierAllCodes = "text-advanced"
	SearchModifierAllCodesIterate      SearchModifierAllCodes = "iterate"
)

// SearchModifierCode represents codes from http://hl7.org/fhir/ValueSet/search-modifier-code
type SearchModifierCode string

const (
	SearchModifierCodeMissing      SearchModifierCode = "missing"
	SearchModifierCodeExact        SearchModifierCode = "exact"
	SearchModifierCodeContains     SearchModifierCode = "contains"
	SearchModifierCodeNot          SearchModifierCode = "not"
	SearchModifierCodeText         SearchModifierCode = "text"
	SearchModifierCodeIn           SearchModifierCode = "in"
	SearchModifierCodeNotIn        SearchModifierCode = "not-in"
	SearchModifierCodeBelow        SearchModifierCode = "below"
	SearchModifierCodeAbove        SearchModifierCode = "above"
	SearchModifierCodeType         SearchModifierCode = "type"
	SearchModifierCodeIdentifier   SearchModifierCode = "identifier"
	SearchModifierCodeOfType       SearchModifierCode = "of-type"
	SearchModifierCodeCodeText     SearchModifierCode = "code-text"
	SearchModifierCodeTextAdvanced SearchModifierCode = "text-advanced"
	SearchModifierCodeIterate      SearchModifierCode = "iterate"
)

// SearchParamType represents codes from http://hl7.org/fhir/ValueSet/search-param-type
type SearchParamType string

const (
	SearchParamTypeNumber    SearchParamType = "number"
	SearchParamTypeDate      SearchParamType = "date"
	SearchParamTypeString    SearchParamType = "string"
	SearchParamTypeToken     SearchParamType = "token"
	SearchParamTypeReference SearchParamType = "reference"
	SearchParamTypeComposite SearchParamType = "composite"
	SearchParamTypeQuantity  SearchParamType = "quantity"
	SearchParamTypeUri       SearchParamType = "uri"
	SearchParamTypeSpecial   SearchParamType = "special"
	SearchParamTypeResource  SearchParamType = "resource"
)

// SearchProcessingModeType represents codes from http://hl7.org/fhir/ValueSet/search-processingmode
type SearchProcessingModeType string

const (
	SearchProcessingModeTypeNormal   SearchProcessingModeType = "normal"
	SearchProcessingModeTypePhonetic SearchProcessingModeType = "phonetic"
	SearchProcessingModeTypeOther    SearchProcessingModeType = "other"
)

// SlicingRules represents codes from http://hl7.org/fhir/ValueSet/resource-slicing-rules
type SlicingRules string

const (
	SlicingRulesClosed    SlicingRules = "closed"
	SlicingRulesOpen      SlicingRules = "open"
	SlicingRulesOpenAtEnd SlicingRules = "openAtEnd"
)

// SlotStatus represents codes from http://hl7.org/fhir/ValueSet/slotstatus
type SlotStatus string

const (
	SlotStatusBusy            SlotStatus = "busy"
	SlotStatusFree            SlotStatus = "free"
	SlotStatusBusyUnavailable SlotStatus = "busy-unavailable"
	SlotStatusBusyTentative   SlotStatus = "busy-tentative"
	SlotStatusEnteredInError  SlotStatus = "entered-in-error"
)

// SortDirection represents codes from http://hl7.org/fhir/ValueSet/sort-direction
type SortDirection string

const (
	SortDirectionAscending  SortDirection = "ascending"
	SortDirectionDescending SortDirection = "descending"
)

// SpecimenCombined represents codes from http://hl7.org/fhir/ValueSet/specimen-combined
type SpecimenCombined string

const (
	SpecimenCombinedGrouped SpecimenCombined = "grouped"
	SpecimenCombinedPooled  SpecimenCombined = "pooled"
)

// SpecimenContainedPreference represents codes from http://hl7.org/fhir/ValueSet/specimen-contained-preference
type SpecimenContainedPreference string

const (
	SpecimenContainedPreferencePreferred SpecimenContainedPreference = "preferred"
	SpecimenContainedPreferenceAlternate SpecimenContainedPreference = "alternate"
)

// SpecimenStatus represents codes from http://hl7.org/fhir/ValueSet/specimen-status
type SpecimenStatus string

const (
	SpecimenStatusAvailable      SpecimenStatus = "available"
	SpecimenStatusUnavailable    SpecimenStatus = "unavailable"
	SpecimenStatusUnsatisfactory SpecimenStatus = "unsatisfactory"
	SpecimenStatusEnteredInError SpecimenStatus = "entered-in-error"
)

// StructureDefinitionKind represents codes from http://hl7.org/fhir/ValueSet/structure-definition-kind
type StructureDefinitionKind string

const (
	StructureDefinitionKindPrimitiveType StructureDefinitionKind = "primitive-type"
	StructureDefinitionKindComplexType   StructureDefinitionKind = "complex-type"
	StructureDefinitionKindResource      StructureDefinitionKind = "resource"
	StructureDefinitionKindLogical       StructureDefinitionKind = "logical"
)

// StructureMapGroupTypeMode represents codes from http://hl7.org/fhir/ValueSet/map-group-type-mode
type StructureMapGroupTypeMode string

const (
	StructureMapGroupTypeModeTypes        StructureMapGroupTypeMode = "types"
	StructureMapGroupTypeModeTypeAndTypes StructureMapGroupTypeMode = "type-and-types"
)

// StructureMapInputMode represents codes from http://hl7.org/fhir/ValueSet/map-input-mode
type StructureMapInputMode string

const (
	StructureMapInputModeSource StructureMapInputMode = "source"
	StructureMapInputModeTarget StructureMapInputMode = "target"
)

// StructureMapModelMode represents codes from http://hl7.org/fhir/ValueSet/map-model-mode
type StructureMapModelMode string

const (
	StructureMapModelModeSource   StructureMapModelMode = "source"
	StructureMapModelModeQueried  StructureMapModelMode = "queried"
	StructureMapModelModeTarget   StructureMapModelMode = "target"
	StructureMapModelModeProduced StructureMapModelMode = "produced"
)

// StructureMapSourceListMode represents codes from http://hl7.org/fhir/ValueSet/map-source-list-mode
type StructureMapSourceListMode string

const (
	StructureMapSourceListModeFirst    StructureMapSourceListMode = "first"
	StructureMapSourceListModeNotFirst StructureMapSourceListMode = "not_first"
	StructureMapSourceListModeLast     StructureMapSourceListMode = "last"
	StructureMapSourceListModeNotLast  StructureMapSourceListMode = "not_last"
	StructureMapSourceListModeOnlyOne  StructureMapSourceListMode = "only_one"
)

// StructureMapTargetListMode represents codes from http://hl7.org/fhir/ValueSet/map-target-list-mode
type StructureMapTargetListMode string

const (
	StructureMapTargetListModeFirst  StructureMapTargetListMode = "first"
	StructureMapTargetListModeShare  StructureMapTargetListMode = "share"
	StructureMapTargetListModeLast   StructureMapTargetListMode = "last"
	StructureMapTargetListModeSingle StructureMapTargetListMode = "single"
)

// StructureMapTransform represents codes from http://hl7.org/fhir/ValueSet/map-transform
type StructureMapTransform string

const (
	StructureMapTransformCreate    StructureMapTransform = "create"
	StructureMapTransformCopy      StructureMapTransform = "copy"
	StructureMapTransformTruncate  StructureMapTransform = "truncate"
	StructureMapTransformEscape    StructureMapTransform = "escape"
	StructureMapTransformCast      StructureMapTransform = "cast"
	StructureMapTransformAppend    StructureMapTransform = "append"
	StructureMapTransformTranslate StructureMapTransform = "translate"
	StructureMapTransformReference StructureMapTransform = "reference"
	StructureMapTransformDateOp    StructureMapTransform = "dateOp"
	StructureMapTransformUuid      StructureMapTransform = "uuid"
	StructureMapTransformPointer   StructureMapTransform = "pointer"
	StructureMapTransformEvaluate  StructureMapTransform = "evaluate"
	StructureMapTransformCc        StructureMapTransform = "cc"
	StructureMapTransformC         StructureMapTransform = "c"
	StructureMapTransformQty       StructureMapTransform = "qty"
	StructureMapTransformId        StructureMapTransform = "id"
	StructureMapTransformCp        StructureMapTransform = "cp"
)

// SubmitDataUpdateType represents codes from http://hl7.org/fhir/ValueSet/submit-data-update-type
type SubmitDataUpdateType string

const (
	SubmitDataUpdateTypeIncremental SubmitDataUpdateType = "incremental"
	SubmitDataUpdateTypeSnapshot    SubmitDataUpdateType = "snapshot"
)

// SubscriptionNotificationType represents codes from http://hl7.org/fhir/ValueSet/subscription-notification-type
type SubscriptionNotificationType string

const (
	SubscriptionNotificationTypeHandshake         SubscriptionNotificationType = "handshake"
	SubscriptionNotificationTypeHeartbeat         SubscriptionNotificationType = "heartbeat"
	SubscriptionNotificationTypeEventNotification SubscriptionNotificationType = "event-notification"
	SubscriptionNotificationTypeQueryStatus       SubscriptionNotificationType = "query-status"
	SubscriptionNotificationTypeQueryEvent        SubscriptionNotificationType = "query-event"
)

// SubscriptionPayloadContent represents codes from http://hl7.org/fhir/ValueSet/subscription-payload-content
type SubscriptionPayloadContent string

const (
	SubscriptionPayloadContentEmpty        SubscriptionPayloadContent = "empty"
	SubscriptionPayloadContentIdOnly       SubscriptionPayloadContent = "id-only"
	SubscriptionPayloadContentFullResource SubscriptionPayloadContent = "full-resource"
)

// SubscriptionStatusCodes represents codes from http://hl7.org/fhir/ValueSet/subscription-status
type SubscriptionStatusCodes string

const (
	SubscriptionStatusCodesRequested      SubscriptionStatusCodes = "requested"
	SubscriptionStatusCodesActive         SubscriptionStatusCodes = "active"
	SubscriptionStatusCodesError          SubscriptionStatusCodes = "error"
	SubscriptionStatusCodesOff            SubscriptionStatusCodes = "off"
	SubscriptionStatusCodesEnteredInError SubscriptionStatusCodes = "entered-in-error"
)

// SupplementedMimeTypes represents codes from http://hl7.org/fhir/ValueSet/supplemented-mimetypes
type SupplementedMimeTypes string

const (
	SupplementedMimeTypesXml  SupplementedMimeTypes = "xml"
	SupplementedMimeTypesJson SupplementedMimeTypes = "json"
	SupplementedMimeTypesTtl  SupplementedMimeTypes = "ttl"
)

// SystemRestfulInteraction represents codes from http://hl7.org/fhir/ValueSet/system-restful-interaction
type SystemRestfulInteraction string

const (
	SystemRestfulInteractionRead                      SystemRestfulInteraction = "read"
	SystemRestfulInteractionVread                     SystemRestfulInteraction = "vread"
	SystemRestfulInteractionUpdate                    SystemRestfulInteraction = "update"
	SystemRestfulInteractionUpdateConditional         SystemRestfulInteraction = "update-conditional"
	SystemRestfulInteractionPatch                     SystemRestfulInteraction = "patch"
	SystemRestfulInteractionPatchConditional          SystemRestfulInteraction = "patch-conditional"
	SystemRestfulInteractionDelete                    SystemRestfulInteraction = "delete"
	SystemRestfulInteractionDeleteConditionalSingle   SystemRestfulInteraction = "delete-conditional-single"
	SystemRestfulInteractionDeleteConditionalMultiple SystemRestfulInteraction = "delete-conditional-multiple"
	SystemRestfulInteractionDeleteHistory             SystemRestfulInteraction = "delete-history"
	SystemRestfulInteractionDeleteHistoryVersion      SystemRestfulInteraction = "delete-history-version"
	SystemRestfulInteractionHistory                   SystemRestfulInteraction = "history"
	SystemRestfulInteractionHistoryInstance           SystemRestfulInteraction = "history-instance"
	SystemRestfulInteractionHistoryType               SystemRestfulInteraction = "history-type"
	SystemRestfulInteractionHistorySystem             SystemRestfulInteraction = "history-system"
	SystemRestfulInteractionCreate                    SystemRestfulInteraction = "create"
	SystemRestfulInteractionCreateConditional         SystemRestfulInteraction = "create-conditional"
	SystemRestfulInteractionSearch                    SystemRestfulInteraction = "search"
	SystemRestfulInteractionSearchType                SystemRestfulInteraction = "search-type"
	SystemRestfulInteractionSearchSystem              SystemRestfulInteraction = "search-system"
	SystemRestfulInteractionSearchCompartment         SystemRestfulInteraction = "search-compartment"
	SystemRestfulInteractionCapabilities              SystemRestfulInteraction = "capabilities"
	SystemRestfulInteractionTransaction               SystemRestfulInteraction = "transaction"
	SystemRestfulInteractionBatch                     SystemRestfulInteraction = "batch"
	SystemRestfulInteractionOperation                 SystemRestfulInteraction = "operation"
)

// TaskIntent represents codes from http://hl7.org/fhir/ValueSet/task-intent
type TaskIntent string

const (
	TaskIntentUnknown       TaskIntent = "unknown"
	TaskIntentProposal      TaskIntent = "proposal"
	TaskIntentSolicitOffer  TaskIntent = "solicit-offer"
	TaskIntentOfferResponse TaskIntent = "offer-response"
	TaskIntentPlan          TaskIntent = "plan"
	TaskIntentDirective     TaskIntent = "directive"
	TaskIntentOrder         TaskIntent = "order"
	TaskIntentOriginalOrder TaskIntent = "original-order"
	TaskIntentReflexOrder   TaskIntent = "reflex-order"
	TaskIntentFillerOrder   TaskIntent = "filler-order"
	TaskIntentInstanceOrder TaskIntent = "instance-order"
	TaskIntentOption        TaskIntent = "option"
)

// TaskStatus represents codes from http://hl7.org/fhir/ValueSet/task-status
type TaskStatus string

const (
	TaskStatusDraft          TaskStatus = "draft"
	TaskStatusRequested      TaskStatus = "requested"
	TaskStatusReceived       TaskStatus = "received"
	TaskStatusAccepted       TaskStatus = "accepted"
	TaskStatusRejected       TaskStatus = "rejected"
	TaskStatusReady          TaskStatus = "ready"
	TaskStatusCancelled      TaskStatus = "cancelled"
	TaskStatusInProgress     TaskStatus = "in-progress"
	TaskStatusOnHold         TaskStatus = "on-hold"
	TaskStatusFailed         TaskStatus = "failed"
	TaskStatusCompleted      TaskStatus = "completed"
	TaskStatusEnteredInError TaskStatus = "entered-in-error"
)

// TriggerType represents codes from http://hl7.org/fhir/ValueSet/trigger-type
type TriggerType string

const (
	TriggerTypeNamedEvent        TriggerType = "named-event"
	TriggerTypePeriodic          TriggerType = "periodic"
	TriggerTypeDataChanged       TriggerType = "data-changed"
	TriggerTypeDataAdded         TriggerType = "data-added"
	TriggerTypeDataModified      TriggerType = "data-modified"
	TriggerTypeDataRemoved       TriggerType = "data-removed"
	TriggerTypeDataAccessed      TriggerType = "data-accessed"
	TriggerTypeDataAccessEnded   TriggerType = "data-access-ended"
	TriggerTypeSubscriptionTopic TriggerType = "subscription-topic"
)

// TriggeredBytype represents codes from http://hl7.org/fhir/ValueSet/observation-triggeredbytype
type TriggeredBytype string

const (
	TriggeredBytypeReflex TriggeredBytype = "reflex"
	TriggeredBytypeRepeat TriggeredBytype = "repeat"
	TriggeredBytypeReRun  TriggeredBytype = "re-run"
)

// TypeDerivationRule represents codes from http://hl7.org/fhir/ValueSet/type-derivation-rule
type TypeDerivationRule string

const (
	TypeDerivationRuleSpecialization TypeDerivationRule = "specialization"
	TypeDerivationRuleConstraint     TypeDerivationRule = "constraint"
)

// TypeRestfulInteraction represents codes from http://hl7.org/fhir/ValueSet/type-restful-interaction
type TypeRestfulInteraction string

const (
	TypeRestfulInteractionRead                      TypeRestfulInteraction = "read"
	TypeRestfulInteractionVread                     TypeRestfulInteraction = "vread"
	TypeRestfulInteractionUpdate                    TypeRestfulInteraction = "update"
	TypeRestfulInteractionUpdateConditional         TypeRestfulInteraction = "update-conditional"
	TypeRestfulInteractionPatch                     TypeRestfulInteraction = "patch"
	TypeRestfulInteractionPatchConditional          TypeRestfulInteraction = "patch-conditional"
	TypeRestfulInteractionDelete                    TypeRestfulInteraction = "delete"
	TypeRestfulInteractionDeleteConditionalSingle   TypeRestfulInteraction = "delete-conditional-single"
	TypeRestfulInteractionDeleteConditionalMultiple TypeRestfulInteraction = "delete-conditional-multiple"
	TypeRestfulInteractionDeleteHistory             TypeRestfulInteraction = "delete-history"
	TypeRestfulInteractionDeleteHistoryVersion      TypeRestfulInteraction = "delete-history-version"
	TypeRestfulInteractionHistory                   TypeRestfulInteraction = "history"
	TypeRestfulInteractionHistoryInstance           TypeRestfulInteraction = "history-instance"
	TypeRestfulInteractionHistoryType               TypeRestfulInteraction = "history-type"
	TypeRestfulInteractionHistorySystem             TypeRestfulInteraction = "history-system"
	TypeRestfulInteractionCreate                    TypeRestfulInteraction = "create"
	TypeRestfulInteractionCreateConditional         TypeRestfulInteraction = "create-conditional"
	TypeRestfulInteractionSearch                    TypeRestfulInteraction = "search"
	TypeRestfulInteractionSearchType                TypeRestfulInteraction = "search-type"
	TypeRestfulInteractionSearchSystem              TypeRestfulInteraction = "search-system"
	TypeRestfulInteractionSearchCompartment         TypeRestfulInteraction = "search-compartment"
	TypeRestfulInteractionCapabilities              TypeRestfulInteraction = "capabilities"
	TypeRestfulInteractionTransaction               TypeRestfulInteraction = "transaction"
	TypeRestfulInteractionBatch                     TypeRestfulInteraction = "batch"
	TypeRestfulInteractionOperation                 TypeRestfulInteraction = "operation"
)

// UDIEntryType represents codes from http://hl7.org/fhir/ValueSet/udi-entry-type
type UDIEntryType string

const (
	UDIEntryTypeBarcode                UDIEntryType = "barcode"
	UDIEntryTypeRfid                   UDIEntryType = "rfid"
	UDIEntryTypeManual                 UDIEntryType = "manual"
	UDIEntryTypeCard                   UDIEntryType = "card"
	UDIEntryTypeSelfReported           UDIEntryType = "self-reported"
	UDIEntryTypeElectronicTransmission UDIEntryType = "electronic-transmission"
	UDIEntryTypeUnknown                UDIEntryType = "unknown"
)

// UnitsOfTime represents codes from http://hl7.org/fhir/ValueSet/units-of-time
type UnitsOfTime string

const (
	UnitsOfTimeS   UnitsOfTime = "s"
	UnitsOfTimeMin UnitsOfTime = "min"
	UnitsOfTimeH   UnitsOfTime = "h"
	UnitsOfTimeD   UnitsOfTime = "d"
	UnitsOfTimeWk  UnitsOfTime = "wk"
	UnitsOfTimeMo  UnitsOfTime = "mo"
	UnitsOfTimeA   UnitsOfTime = "a"
)

// Use represents codes from http://hl7.org/fhir/ValueSet/claim-use
type Use string

const (
	UseClaim            Use = "claim"
	UsePreauthorization Use = "preauthorization"
	UsePredetermination Use = "predetermination"
)

// ValueFilterComparator represents codes from http://hl7.org/fhir/ValueSet/value-filter-comparator
type ValueFilterComparator string

const (
	ValueFilterComparatorEq ValueFilterComparator = "eq"
	ValueFilterComparatorNe ValueFilterComparator = "ne"
	ValueFilterComparatorGt ValueFilterComparator = "gt"
	ValueFilterComparatorLt ValueFilterComparator = "lt"
	ValueFilterComparatorGe ValueFilterComparator = "ge"
	ValueFilterComparatorLe ValueFilterComparator = "le"
	ValueFilterComparatorSa ValueFilterComparator = "sa"
	ValueFilterComparatorEb ValueFilterComparator = "eb"
	ValueFilterComparatorAp ValueFilterComparator = "ap"
)

// VersionIndependentResourceTypesAll represents codes from http://hl7.org/fhir/ValueSet/version-independent-all-resource-types
type VersionIndependentResourceTypesAll string

const (
	VersionIndependentResourceTypesAllBodySite                          VersionIndependentResourceTypesAll = "BodySite"
	VersionIndependentResourceTypesAllCatalogEntry                      VersionIndependentResourceTypesAll = "CatalogEntry"
	VersionIndependentResourceTypesAllConformance                       VersionIndependentResourceTypesAll = "Conformance"
	VersionIndependentResourceTypesAllDataElement                       VersionIndependentResourceTypesAll = "DataElement"
	VersionIndependentResourceTypesAllDeviceComponent                   VersionIndependentResourceTypesAll = "DeviceComponent"
	VersionIndependentResourceTypesAllDeviceUseRequest                  VersionIndependentResourceTypesAll = "DeviceUseRequest"
	VersionIndependentResourceTypesAllDeviceUseStatement                VersionIndependentResourceTypesAll = "DeviceUseStatement"
	VersionIndependentResourceTypesAllDiagnosticOrder                   VersionIndependentResourceTypesAll = "DiagnosticOrder"
	VersionIndependentResourceTypesAllDocumentManifest                  VersionIndependentResourceTypesAll = "DocumentManifest"
	VersionIndependentResourceTypesAllEffectEvidenceSynthesis           VersionIndependentResourceTypesAll = "EffectEvidenceSynthesis"
	VersionIndependentResourceTypesAllEligibilityRequest                VersionIndependentResourceTypesAll = "EligibilityRequest"
	VersionIndependentResourceTypesAllEligibilityResponse               VersionIndependentResourceTypesAll = "EligibilityResponse"
	VersionIndependentResourceTypesAllExpansionProfile                  VersionIndependentResourceTypesAll = "ExpansionProfile"
	VersionIndependentResourceTypesAllImagingManifest                   VersionIndependentResourceTypesAll = "ImagingManifest"
	VersionIndependentResourceTypesAllImagingObjectSelection            VersionIndependentResourceTypesAll = "ImagingObjectSelection"
	VersionIndependentResourceTypesAllMedia                             VersionIndependentResourceTypesAll = "Media"
	VersionIndependentResourceTypesAllMedicationOrder                   VersionIndependentResourceTypesAll = "MedicationOrder"
	VersionIndependentResourceTypesAllMedicationUsage                   VersionIndependentResourceTypesAll = "MedicationUsage"
	VersionIndependentResourceTypesAllMedicinalProduct                  VersionIndependentResourceTypesAll = "MedicinalProduct"
	VersionIndependentResourceTypesAllMedicinalProductAuthorization     VersionIndependentResourceTypesAll = "MedicinalProductAuthorization"
	VersionIndependentResourceTypesAllMedicinalProductContraindication  VersionIndependentResourceTypesAll = "MedicinalProductContraindication"
	VersionIndependentResourceTypesAllMedicinalProductIndication        VersionIndependentResourceTypesAll = "MedicinalProductIndication"
	VersionIndependentResourceTypesAllMedicinalProductIngredient        VersionIndependentResourceTypesAll = "MedicinalProductIngredient"
	VersionIndependentResourceTypesAllMedicinalProductInteraction       VersionIndependentResourceTypesAll = "MedicinalProductInteraction"
	VersionIndependentResourceTypesAllMedicinalProductManufactured      VersionIndependentResourceTypesAll = "MedicinalProductManufactured"
	VersionIndependentResourceTypesAllMedicinalProductPackaged          VersionIndependentResourceTypesAll = "MedicinalProductPackaged"
	VersionIndependentResourceTypesAllMedicinalProductPharmaceutical    VersionIndependentResourceTypesAll = "MedicinalProductPharmaceutical"
	VersionIndependentResourceTypesAllMedicinalProductUndesirableEffect VersionIndependentResourceTypesAll = "MedicinalProductUndesirableEffect"
	VersionIndependentResourceTypesAllOrder                             VersionIndependentResourceTypesAll = "Order"
	VersionIndependentResourceTypesAllOrderResponse                     VersionIndependentResourceTypesAll = "OrderResponse"
	VersionIndependentResourceTypesAllProcedureRequest                  VersionIndependentResourceTypesAll = "ProcedureRequest"
	VersionIndependentResourceTypesAllProcessRequest                    VersionIndependentResourceTypesAll = "ProcessRequest"
	VersionIndependentResourceTypesAllProcessResponse                   VersionIndependentResourceTypesAll = "ProcessResponse"
	VersionIndependentResourceTypesAllReferralRequest                   VersionIndependentResourceTypesAll = "ReferralRequest"
	VersionIndependentResourceTypesAllRequestGroup                      VersionIndependentResourceTypesAll = "RequestGroup"
	VersionIndependentResourceTypesAllResearchDefinition                VersionIndependentResourceTypesAll = "ResearchDefinition"
	VersionIndependentResourceTypesAllResearchElementDefinition         VersionIndependentResourceTypesAll = "ResearchElementDefinition"
	VersionIndependentResourceTypesAllRiskEvidenceSynthesis             VersionIndependentResourceTypesAll = "RiskEvidenceSynthesis"
	VersionIndependentResourceTypesAllSequence                          VersionIndependentResourceTypesAll = "Sequence"
	VersionIndependentResourceTypesAllServiceDefinition                 VersionIndependentResourceTypesAll = "ServiceDefinition"
	VersionIndependentResourceTypesAllSubstanceSpecification            VersionIndependentResourceTypesAll = "SubstanceSpecification"
)

// VisionBase represents codes from http://hl7.org/fhir/ValueSet/vision-base-codes
type VisionBase string

const (
	VisionBaseUp   VisionBase = "up"
	VisionBaseDown VisionBase = "down"
	VisionBaseIn   VisionBase = "in"
	VisionBaseOut  VisionBase = "out"
)

// VisionEyes represents codes from http://hl7.org/fhir/ValueSet/vision-eye-codes
type VisionEyes string

const (
	VisionEyesRight VisionEyes = "right"
	VisionEyesLeft  VisionEyes = "left"
)

// WeekOfMonth represents codes from http://hl7.org/fhir/ValueSet/week-of-month
type WeekOfMonth string

const (
	WeekOfMonthFirst  WeekOfMonth = "first"
	WeekOfMonthSecond WeekOfMonth = "second"
	WeekOfMonthThird  WeekOfMonth = "third"
	WeekOfMonthFourth WeekOfMonth = "fourth"
	WeekOfMonthLast   WeekOfMonth = "last"
)
