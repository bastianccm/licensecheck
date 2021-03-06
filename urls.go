// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package licensecheck

// This file contains the list of known URLs for valid licenses, only.
// Do not add other code.

// TODO: Find a canonical source of this information, or at least a
// disciplined way to develop it. This is cadged from gopkg.in/src-d/go-license-detector.v2.

// The map goes from URL to name. The URL text does not include the leading http:// or https://
// or trailing /.
// All entries are lower case.
// Keep this list sorted for easy checking.
var urlLicenses = map[string]string{
	"creativecommons.org/licenses/by-nc-nd/2.0":                                "CC-BY-NC-ND-2.0",
	"creativecommons.org/licenses/by-nc-nd/2.5":                                "CC-BY-NC-ND-2.5",
	"creativecommons.org/licenses/by-nc-nd/3.0":                                "CC-BY-NC-ND-3.0",
	"creativecommons.org/licenses/by-nc-nd/4.0":                                "CC-BY-NC-ND-4.0",
	"creativecommons.org/licenses/by-nc-sa/1.0":                                "CC-BY-NC-SA-1.0",
	"creativecommons.org/licenses/by-nc-sa/2.0":                                "CC-BY-NC-SA-2.0",
	"creativecommons.org/licenses/by-nc-sa/2.5":                                "CC-BY-NC-SA-2.5",
	"creativecommons.org/licenses/by-nc-sa/3.0":                                "CC-BY-NC-SA-3.0",
	"creativecommons.org/licenses/by-nc-sa/4.0":                                "CC-BY-NC-SA-4.0",
	"creativecommons.org/licenses/by-nc/1.0":                                   "CC-BY-NC-1.0",
	"creativecommons.org/licenses/by-nc/2.0":                                   "CC-BY-NC-2.0",
	"creativecommons.org/licenses/by-nc/2.5":                                   "CC-BY-NC-2.5",
	"creativecommons.org/licenses/by-nc/3.0":                                   "CC-BY-NC-3.0",
	"creativecommons.org/licenses/by-nc/4.0":                                   "CC-BY-NC-4.0",
	"creativecommons.org/licenses/by-nd-nc/1.0":                                "CC-BY-NC-ND-1.0",
	"creativecommons.org/licenses/by-nd/1.0":                                   "CC-BY-ND-1.0",
	"creativecommons.org/licenses/by-nd/2.0":                                   "CC-BY-ND-2.0",
	"creativecommons.org/licenses/by-nd/2.5":                                   "CC-BY-ND-2.5",
	"creativecommons.org/licenses/by-nd/3.0":                                   "CC-BY-ND-3.0",
	"creativecommons.org/licenses/by-nd/4.0":                                   "CC-BY-ND-4.0",
	"creativecommons.org/licenses/by-sa/1.0":                                   "CC-BY-SA-1.0",
	"creativecommons.org/licenses/by-sa/2.0":                                   "CC-BY-SA-2.0",
	"creativecommons.org/licenses/by-sa/2.5":                                   "CC-BY-SA-2.5",
	"creativecommons.org/licenses/by-sa/3.0":                                   "CC-BY-SA-3.0",
	"creativecommons.org/licenses/by-sa/4.0":                                   "CC-BY-SA-4.0",
	"creativecommons.org/licenses/by/1.0":                                      "CC-BY-1.0",
	"creativecommons.org/licenses/by/2.0":                                      "CC-BY-2.0",
	"creativecommons.org/licenses/by/2.5":                                      "CC-BY-2.5",
	"creativecommons.org/licenses/by/3.0":                                      "CC-BY-3.0",
	"creativecommons.org/licenses/by/4.0":                                      "CC-BY-4.0",
	"creativecommons.org/publicdomain/zero/1.0":                                "CC0-1.0",
	"opensource.org/licenses/apache-1.1":                                       "Apache-1.1",
	"opensource.org/licenses/artistic-1.0":                                     "Artistic-1.0",
	"opensource.org/licenses/bsdpluspatent":                                    "BSD-2-Clause-Patent",
	"opensource.org/licenses/catosl-1.1":                                       "CATOSL-1.1",
	"opensource.org/licenses/cpl-1.0":                                          "CPL-1.0",
	"opensource.org/licenses/cua-opl-1.0":                                      "CUA-OPL-1.0",
	"opensource.org/licenses/ecl-1.0":                                          "ECL-1.0",
	"opensource.org/licenses/ecl-2.0":                                          "ECL-2.0",
	"opensource.org/licenses/efl-1.0":                                          "EFL-1.0",
	"opensource.org/licenses/efl-2.0":                                          "EFL-2.0",
	"opensource.org/licenses/entessa":                                          "Entessa",
	"opensource.org/licenses/intel":                                            "Intel",
	"opensource.org/licenses/lpl-1.0":                                          "LPL-1.0",
	"opensource.org/licenses/liliq-p-1.1":                                      "LiLiQ-P-1.1",
	"opensource.org/licenses/liliq-r-1.1":                                      "LiLiQ-R-1.1",
	"opensource.org/licenses/liliq-rplus-1.1":                                  "LiLiQ-Rplus-1.1",
	"opensource.org/licenses/mpl-1.0":                                          "MPL-1.0",
	"opensource.org/licenses/mpl-2.0":                                          "MPL-2.0",
	"opensource.org/licenses/opl-2.1":                                          "OSET-PL-2.1",
	"opensource.org/licenses/osl-1.0":                                          "OSL-1.0",
	"opensource.org/licenses/osl-2.1":                                          "OSL-2.1",
	"opensource.org/licenses/rpl-1.1":                                          "RPL-1.1",
	"opensource.org/licenses/sissl":                                            "SISSL",
	"opensource.org/licenses/upl":                                              "UPL-1.0",
	"opensource.org/licenses/xnet":                                             "Xnet",
	"opensource.org/licenses/zpl-2.0":                                          "ZPL-2.0",
	"www.apache.org/licenses/license-2.0":                                      "Apache-2.0",
	"www.gnu.org/licenses/agpl.txt":                                            "AGPL-3.0",
	"www.gnu.org/licenses/autoconf-exception-3.0.html":                         "GPL-3.0-with-autoconf-exception",
	"www.gnu.org/licenses/ecos-license.html":                                   "eCos-2.0",
	"www.gnu.org/licenses/fdl-1.3.txt":                                         "GFDL-1.3",
	"www.gnu.org/licenses/gcc-exception-3.1.html":                              "GPL-3.0-with-GCC-exception",
	"www.gnu.org/licenses/gpl-3.0-standalone.html":                             "GPL-3.0",
	"www.gnu.org/licenses/gpl-faq.html#fontexception":                          "GPL-2.0-with-font-exception",
	"www.gnu.org/licenses/lgpl-3.0-standalone.html":                            "LGPL-3.0",
	"www.gnu.org/licenses/old-licenses/fdl-1.1.txt":                            "GFDL-1.1",
	"www.gnu.org/licenses/old-licenses/gpl-1.0-standalone.html":                "GPL-1.0",
	"www.gnu.org/licenses/old-licenses/gpl-2.0-standalone.html":                "GPL-2.0",
	"www.gnu.org/licenses/old-licenses/lgpl-2.0-standalone.html":               "LGPL-2.0",
	"www.gnu.org/licenses/old-licenses/lgpl-2.1-standalone.html":               "LGPL-2.1",
	"www.gnu.org/prep/maintain/html_node/license-notices-for-other-files.html": "FSFAP",
	"www.gnu.org/software/classpath/license.html":                              "GPL-2.0-with-classpath-exception",
	"www.opensource.org/licenses/agpl-3.0":                                     "AGPL-3.0",
	"www.opensource.org/licenses/apl-1.0":                                      "APL-1.0",
	"www.opensource.org/licenses/apache-2.0":                                   "Apache-2.0",
	"www.opensource.org/licenses/bsd-2-clause":                                 "BSD-2-Clause",
	"www.opensource.org/licenses/bsd-3-clause":                                 "BSD-3-Clause",
	"www.opensource.org/licenses/bsl-1.0":                                      "BSL-1.0",
	"www.opensource.org/licenses/cnri-python":                                  "CNRI-Python",
	"www.opensource.org/licenses/cpal-1.0":                                     "CPAL-1.0",
	"www.opensource.org/licenses/epl-1.0":                                      "EPL-1.0",
	"www.opensource.org/licenses/epl-2.0":                                      "EPL-2.0",
	"www.opensource.org/licenses/eudatagrid":                                   "EUDatagrid",
	"www.opensource.org/licenses/eupl-1.1":                                     "EUPL-1.1",
	"www.opensource.org/licenses/fair":                                         "Fair",
	"www.opensource.org/licenses/frameworx-1.0":                                "Frameworx-1.0",
	"www.opensource.org/licenses/gpl-2.0":                                      "GPL-2.0",
	"www.opensource.org/licenses/gpl-3.0":                                      "GPL-3.0",
	"www.opensource.org/licenses/hpnd":                                         "HPND",
	"www.opensource.org/licenses/ipa":                                          "IPA",
	"www.opensource.org/licenses/ipl-1.0":                                      "IPL-1.0",
	"www.opensource.org/licenses/isc":                                          "ISC",
	"www.opensource.org/licenses/lgpl-2.1":                                     "LGPL-2.1",
	"www.opensource.org/licenses/lgpl-3.0":                                     "LGPL-3.0",
	"www.opensource.org/licenses/lpl-1.02":                                     "LPL-1.02",
	"www.opensource.org/licenses/lppl-1.3c":                                    "LPPL-1.3c",
	"www.opensource.org/licenses/mit":                                          "MIT",
	"www.opensource.org/licenses/mpl-1.1":                                      "MPL-1.1",
	"www.opensource.org/licenses/ms-pl":                                        "MS-PL",
	"www.opensource.org/licenses/ms-rl":                                        "MS-RL",
	"www.opensource.org/licenses/miros":                                        "MirOS",
	"www.opensource.org/licenses/motosoto":                                     "Motosoto",
	"www.opensource.org/licenses/multics":                                      "Multics",
	"www.opensource.org/licenses/nasa-1.3":                                     "NASA-1.3",
	"www.opensource.org/licenses/ncsa":                                         "NCSA",
	"www.opensource.org/licenses/ngpl":                                         "NGPL",
	"www.opensource.org/licenses/nosl3.0":                                      "NPOSL-3.0",
	"www.opensource.org/licenses/ntp":                                          "NTP",
	"www.opensource.org/licenses/naumen":                                       "Naumen",
	"www.opensource.org/licenses/oclc-2.0":                                     "OCLC-2.0",
	"www.opensource.org/licenses/ofl-1.1":                                      "OFL-1.1",
	"www.opensource.org/licenses/ogtsl":                                        "OGTSL",
	"www.opensource.org/licenses/osl-3.0":                                      "OSL-3.0",
	"www.opensource.org/licenses/php-3.0":                                      "PHP-3.0",
	"www.opensource.org/licenses/postgresql":                                   "PostgreSQL",
	"www.opensource.org/licenses/python-2.0":                                   "Python-2.0",
	"www.opensource.org/licenses/qpl-1.0":                                      "QPL-1.0",
	"www.opensource.org/licenses/rpl-1.5":                                      "RPL-1.5",
	"www.opensource.org/licenses/rpsl-1.0":                                     "RPSL-1.0",
	"www.opensource.org/licenses/rscpl":                                        "RSCPL",
	"www.opensource.org/licenses/spl-1.0":                                      "SPL-1.0",
	"www.opensource.org/licenses/simpl-2.0":                                    "SimPL-2.0",
	"www.opensource.org/licenses/sleepycat":                                    "Sleepycat",
	"www.opensource.org/licenses/vsl-1.0":                                      "VSL-1.0",
	"www.opensource.org/licenses/w3c":                                          "W3C",
	"www.opensource.org/licenses/wxwindows":                                    "wxWindows",
	"www.opensource.org/licenses/watcom-1.0":                                   "Watcom-1.0",
	"www.opensource.org/licenses/zlib":                                         "Zlib",
	"www.opensource.org/licenses/afl-3.0":                                      "AFL-3.0",
	"www.opensource.org/licenses/artistic-license-2.0":                         "Artistic-2.0",
	"www.opensource.org/licenses/attribution":                                  "AAL",
	"www.opensource.org/licenses/cddl1":                                        "CDDL-1.0",
	"www.opensource.org/licenses/nokia":                                        "Nokia",
}
