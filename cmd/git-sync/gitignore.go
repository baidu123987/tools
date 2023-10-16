package main

import "github.com/gobwas/glob"

// Taken from cdnjs/cdnjs' .gitignore file
var (
	GIT_IGNORE = []glob.Glob{
		// folders need to be ingored due to auto auto-update mechanism
		glob.MustCompile("/ajax/libs/agility/0.1.0-docs/"),
		glob.MustCompile("/ajax/libs/angular-bootstrap-colorpicker/push/"),
		glob.MustCompile("/ajax/libs/angular-chart.js/rm/"),
		glob.MustCompile("/ajax/libs/angular-gantt/.0.6.1"),
		glob.MustCompile("/ajax/libs/angular-google-maps/2.0.0-SNAPSHOT*"),
		glob.MustCompile("/ajax/libs/angular-gridster/v.0.10.4"),
		glob.MustCompile("/ajax/libs/angular-leaflet-directive/delete/"),
		glob.MustCompile("/ajax/libs/angular-leaflet-directive/remove/"),
		glob.MustCompile("/ajax/libs/angular-material/0.0.2-6071305/"),
		glob.MustCompile("/ajax/libs/angular-material/0.0.2-ca7793d/"),
		glob.MustCompile("/ajax/libs/angular-material/0.0.2-e59e064cd0e19d841e9c252f973a9eb7e87543a2/"),
		glob.MustCompile("/ajax/libs/angular-material/*-master-*"),
		glob.MustCompile("/ajax/libs/angular-mocks/*-build.*/"),
		glob.MustCompile("/ajax/libs/angular-smart-table/help/"),
		glob.MustCompile("/ajax/libs/angulartics/googleanalytics/"),
		glob.MustCompile("/ajax/libs/angular-ui-router/*-SNAPSHOT*"),
		glob.MustCompile("/ajax/libs/appbase-js/1.0.0/"),
		glob.MustCompile("/ajax/libs/ant-design-pro/2.3.0/"),
		glob.MustCompile("/ajax/libs/atmosphere/javascript-project-*"),
		glob.MustCompile("/ajax/libs/aui/*do-not-use/"),
		glob.MustCompile("/ajax/libs/aui/*SNAPSHOT*/"),
		glob.MustCompile("/ajax/libs/avalon.js/2.1.16ee/"),
		glob.MustCompile("/ajax/libs/avalon.js/before-rebuild-strategy/"),
		glob.MustCompile("/ajax/libs/avalon.js/dynamic_detect/"),
		glob.MustCompile("/ajax/libs/avalon.js/ms_controller_fix/"),
		glob.MustCompile("/ajax/libs/aws-sdk/*/appinfo.sample.js"),
		glob.MustCompile("/ajax/libs/axe-core/*-canary.*"),
		glob.MustCompile("/ajax/libs/ag-grid/vg1tisc00l/"),
		glob.MustCompile("/ajax/libs/aspnet-signalr/*-rc*"),
		glob.MustCompile("/ajax/libs/aspnet-signalr/*-preview*"),
		glob.MustCompile("/ajax/libs/browser-logos/[a-z]*"),
		glob.MustCompile("/ajax/libs/backbone.layoutmanager/0.0.0/"),
		glob.MustCompile("/ajax/libs/backbone-relational/0.8.0plus/"),
		glob.MustCompile("/ajax/libs/bean/1.0.0a1/"),
		glob.MustCompile("/ajax/libs/bonzo/1.3.2-clone/"),
		glob.MustCompile("/ajax/libs/bokeh/1.3.2-dev2/"),
		glob.MustCompile("/ajax/libs/bPopup/*b*"),
		glob.MustCompile("/ajax/libs/bPopup/ver*"),
		glob.MustCompile("/ajax/libs/brython/3.2.*"),
		glob.MustCompile("/ajax/libs/bxslider/*/jquery-3.1.1.min.js"),
		glob.MustCompile("/ajax/libs/clappr/media_control_youtube/"),
		glob.MustCompile("/ajax/libs/clarity-icons/*-dev*"),
		glob.MustCompile("/ajax/libs/clone/0.1.*"),
		glob.MustCompile("/ajax/libs/clusterize.js/v.0.6.0"),
		glob.MustCompile("/ajax/libs/codemirror/3.11.01/"),
		glob.MustCompile("/ajax/libs/codemirror/3.12.00/"),
		glob.MustCompile("/ajax/libs/codemirror/3.13.00/"),
		glob.MustCompile("/ajax/libs/condition/1.0.1b/"),
		glob.MustCompile("/ajax/libs/core-js/1,0,1/"),
		glob.MustCompile("/ajax/libs/crypto-js/release-3.1.2*"),
		glob.MustCompile("/ajax/libs/d3-tip/0.7.0-temp-fork/"),
		glob.MustCompile("/ajax/libs/datatables/BETA*"),
		glob.MustCompile("/ajax/libs/datatables/RELEASE*"),
		glob.MustCompile("/ajax/libs/date-fns/v.*"),
		glob.MustCompile("/ajax/libs/deck.js/latest/"),
		glob.MustCompile("/ajax/libs/design-system/2.6.0-dev"),
		glob.MustCompile("/ajax/libs/device.js/0.1.16/"),
		glob.MustCompile("/ajax/libs/diva.js/RELEASE-*/"),
		glob.MustCompile("/ajax/libs/dygraph/1.0.1-bower"),
		glob.MustCompile("/ajax/libs/dygraph/1.0.1+bower"),
		glob.MustCompile("/ajax/libs/elfinder/*-src"),
		glob.MustCompile("/ajax/libs/embed-js/*-alpha.*"),
		glob.MustCompile("/ajax/libs/esri-leaflet-geocoder/2.2.5"),
		glob.MustCompile("/ajax/libs/europa/4.0.0-alpha*/"),
		glob.MustCompile("/ajax/libs/fixed-data-table/latest/"),
		glob.MustCompile("/ajax/libs/flexslider/version"),
		glob.MustCompile("/ajax/libs/flocks.js/0.15.7/"),
		glob.MustCompile("/ajax/libs/foggy/foggy*"),
		glob.MustCompile("/ajax/libs/formstone/remove/"),
		glob.MustCompile("/ajax/libs/foundation/undefined/"),
		glob.MustCompile("/ajax/libs/fpo/[3-5].*"),
		glob.MustCompile("/ajax/libs/froala-editor/froala_editor_1.0.3/"),
		glob.MustCompile("/ajax/libs/froala-editor/v.1.0.5/"),
		glob.MustCompile("/ajax/libs/fullPage.js/v.2.0.1"),
		glob.MustCompile("/ajax/libs/fullPage.js/v.2.2.8"),
		glob.MustCompile("/ajax/libs/fullPage.js/v.2.7.5"),
		glob.MustCompile("/ajax/libs/gemma/0.1.0/"),
		glob.MustCompile("/ajax/libs/gemma/0.4.0/"),
		glob.MustCompile("/ajax/libs/graphdracula/1.0.1-*"),
		glob.MustCompile("/ajax/libs/guards/bootstrap-*/"),
		glob.MustCompile("/ajax/libs/guards/foundation-*/"),
		glob.MustCompile("/ajax/libs/hls.js/*-canary.*/"),
		glob.MustCompile("/ajax/libs/Hyphenator/HyphenatorBeta12/"),
		glob.MustCompile("/ajax/libs/Hyphenator/Version_*"),
		glob.MustCompile("/ajax/libs/imgareaselect/0.9.11-rc.1/"),
		glob.MustCompile("/ajax/libs/inferno/*-alpha.*"),
		glob.MustCompile("/ajax/libs/inferno-compat/*-alpha.*/"),
		glob.MustCompile("/ajax/libs/inferno-component/*-alpha.*"),
		glob.MustCompile("/ajax/libs/inferno-create-class/*-alpha.*"),
		glob.MustCompile("/ajax/libs/inferno-create-element/*-alpha.*"),
		glob.MustCompile("/ajax/libs/inferno-devtools/*-alpha.*"),
		glob.MustCompile("/ajax/libs/inferno-hyperscript/*-alpha.*"),
		glob.MustCompile("/ajax/libs/inferno-mobx/*-alpha.*"),
		glob.MustCompile("/ajax/libs/inferno-redux/*-alpha.*"),
		glob.MustCompile("/ajax/libs/inferno-router/*-alpha.*"),
		glob.MustCompile("/ajax/libs/inferno-server/*-alpha.*"),
		glob.MustCompile("/ajax/libs/inferno-shared/*-alpha.*"),
		glob.MustCompile("/ajax/libs/inferno-test-utils/*-alpha.*"),
		glob.MustCompile("/ajax/libs/inferno-vnode-flags/*-alpha.*"),
		glob.MustCompile("/ajax/libs/infusion/*-dev.*/"),
		glob.MustCompile("/ajax/libs/interactive-data-display/1.0.1-DomTracking"),
		glob.MustCompile("/ajax/libs/interactive-data-display/1.0.1-DomTracking"),
		glob.MustCompile("/ajax/libs/interactive-data-display/*build*"),
		glob.MustCompile("/ajax/libs/ionic-framework/2.0.0-alpha.2.0.0-alpha.50/"),
		glob.MustCompile("/ajax/libs/ionicons/3.0.0-alpha.1/"),
		glob.MustCompile("/ajax/libs/ionic/testing"),
		glob.MustCompile("/ajax/libs/jarallax/1.6.1_/"),
		glob.MustCompile("/ajax/libs/jasmine/list/"),
		glob.MustCompile("/ajax/libs/jqBootstrapValidation/2.0.0-alpha"),
		glob.MustCompile("/ajax/libs/jQRangeSlider/latest/"),
		glob.MustCompile("/ajax/libs/jquery.alphanum/Stable_1.0"),
		glob.MustCompile("/ajax/libs/jquery.alphanum/Stable_1.1"),
		glob.MustCompile("/ajax/libs/jquery-bracket/release-*"),
		glob.MustCompile("/ajax/libs/jquery-countdown/Release_1*"),
		glob.MustCompile("/ajax/libs/jquery-datetimepicker/v.1.0.1"),
		glob.MustCompile("/ajax/libs/jquery-easing/1.3.0/"),
		glob.MustCompile("/ajax/libs/jquery-easing/1.3.1/"),
		glob.MustCompile("/ajax/libs/jquery-form-validator/2.2"),
		glob.MustCompile("/ajax/libs/jquery.gridster/v.0.4.0/"),
		glob.MustCompile("/ajax/libs/jquery.ime/1.0.0+20131229"),
		glob.MustCompile("/ajax/libs/jquery-tokeninput/jquery-tokeninput-1.*"),
		glob.MustCompile("/ajax/libs/jquery-ui-timepicker-addon/0.0.0-ignored/"),
		glob.MustCompile("/ajax/libs/jsmediatags/.3.2.0/"),
		glob.MustCompile("/ajax/libs/jsmediatags/delete/"),
		glob.MustCompile("/ajax/libs/js-nacl/last_plain_nacl_version/"),
		glob.MustCompile("/ajax/libs/jspdf/v.1.4.0/"),
		glob.MustCompile("/ajax/libs/jsrender/*/jsrender-node.js"),
		glob.MustCompile("/ajax/libs/js-scrypt/1.1.0/"),
		glob.MustCompile("/ajax/libs/jsSHA/release*"),
		glob.MustCompile("/ajax/libs/jsxgraph/RELEASE*"),
		glob.MustCompile("/ajax/libs/jsxgraph/rm/"),
		glob.MustCompile("/ajax/libs/jsxgraph/sketchometry*"),
		glob.MustCompile("/ajax/libs/jsxgraph/skm*"),
		glob.MustCompile("/ajax/libs/juicer/0.5.*-stable"),
		glob.MustCompile("/ajax/libs/juicer/0.6.*-stable*"),
		glob.MustCompile("/ajax/libs/knockout-sortable/v.0.14.0/"),
		glob.MustCompile("/ajax/libs/libphonenumber/libphonenumber*/"),
		glob.MustCompile("/ajax/libs/loadCSS/2.0.0-[2-4]"),
		glob.MustCompile("/ajax/libs/LoadGo/1.0/"),
		glob.MustCompile("/ajax/libs/lodash.js/*-es"),
		glob.MustCompile("/ajax/libs/lodash.js/*-npm"),
		glob.MustCompile("/ajax/libs/material-design-icons/2.0"),
		glob.MustCompile("/ajax/libs/material-design-icons/*/gulpfile.babel.js"),
		glob.MustCompile("/ajax/libs/mdbootstrap/4.4.0-dev"),
		glob.MustCompile("/ajax/libs/meshki/1.0.0/"),
		glob.MustCompile("/ajax/libs/meshki/1.1.0/"),
		glob.MustCompile("/ajax/libs/messenger/.1.5.0/"),
		glob.MustCompile("/ajax/libs/metrics-graphics/v.7.0"),
		glob.MustCompile("/ajax/libs/min/1.5"),
		glob.MustCompile("/ajax/libs/mini.css/3.0.0-alpha.3"),
		glob.MustCompile("/ajax/libs/mobx/*albert*"),
		glob.MustCompile("/ajax/libs/mobx/*fix*"),
		glob.MustCompile("/ajax/libs/mobx-react/*-fix*/"),
		glob.MustCompile("/ajax/libs/mobx-react/*-observableprops*/"),
		glob.MustCompile("/ajax/libs/mojio-js/connery/"),
		glob.MustCompile("/ajax/libs/mojio-js/oauth-token/"),
		glob.MustCompile("/ajax/libs/mojio-js/robb/"),
		glob.MustCompile("/ajax/libs/msl-client-browser/msl-client-java-1.0.0/"),
		glob.MustCompile("/ajax/libs/musicmetadata/Browser-0.4.3/"),
		glob.MustCompile("/ajax/libs/ng-pdfviewer/r0.1.0/"),
		glob.MustCompile("/ajax/libs/ng-pdfviewer/r0.2.0/"),
		glob.MustCompile("/ajax/libs/ng-pdfviewer/r0.2.1/"),
		glob.MustCompile("/ajax/libs/nuclear-js/show"),
		glob.MustCompile("/ajax/libs/numbered/0.0.4"),
		glob.MustCompile("/ajax/libs/numbered/1.0.0"),
		glob.MustCompile("/ajax/libs/nunjucks/v.1.3.1"),
		glob.MustCompile("/ajax/libs/octicons/0.0.0-*"),
		glob.MustCompile("/ajax/libs/octicons/*-alpha.*"),
		glob.MustCompile("/ajax/libs/octicons/*-rc.*"),
		glob.MustCompile("/ajax/libs/onsen/0.6.0/"),
		glob.MustCompile("/ajax/libs/onsen/2.10.7/"),
		glob.MustCompile("/ajax/libs/orb/1.0.0/"),
		glob.MustCompile("/ajax/libs/origamijs/0.0.2/"),
		glob.MustCompile("/ajax/libs/paper.js/src/"),
		glob.MustCompile("/ajax/libs/PhysicsJS/physicsjs-*"),
		glob.MustCompile("/ajax/libs/plottable.js/3.0.0-*-g*"),
		glob.MustCompile("/ajax/libs/plottable.js/*-dev-*"),
		glob.MustCompile("/ajax/libs/plottable.js/*-hover-tooltip*"),
		glob.MustCompile("/ajax/libs/plottable.js/v.0.18.0"),
		glob.MustCompile("/ajax/libs/plottable.js/v.2.4.1"),
		glob.MustCompile("/ajax/libs/postgrest-client/1.0.*"),
		glob.MustCompile("/ajax/libs/postscribe/ersion"),
		glob.MustCompile("/ajax/libs/prettyCheckable/v.2.1.1/"),
		glob.MustCompile("/ajax/libs/Primer/*-alpha*"),
		glob.MustCompile("/ajax/libs/punycode/2.0.*"),
		glob.MustCompile("/ajax/libs/punycode/2.1.*"),
		glob.MustCompile("/ajax/libs/qoopido.js/support"),
		glob.MustCompile("/ajax/libs/qoopido.js/widget-adapt"),
		glob.MustCompile("/ajax/libs/randomcolor/0.2"),
		glob.MustCompile("/ajax/libs/rangy/1.3.0-alpha.*"),
		glob.MustCompile("/ajax/libs/rangy/1.3.0-beta.*"),
		glob.MustCompile("/ajax/libs/raphael/semver/"),
		glob.MustCompile("/ajax/libs/react/0.0.0-*/"),
		glob.MustCompile("/ajax/libs/react/*-alpha.*[a-z]*/"),
		glob.MustCompile("/ajax/libs/react-data-grid/*-*/"),
		glob.MustCompile("/ajax/libs/react-data-grid/*-alpha*/"),
		glob.MustCompile("/ajax/libs/react-data-grid/*-dm*/"),
		glob.MustCompile("/ajax/libs/react-dom/0.0.0-*/"),
		glob.MustCompile("/ajax/libs/react-dom/*-alpha.*[a-z]*/"),
		glob.MustCompile("/ajax/libs/react-grid-layout/0.[0-7].*"),
		glob.MustCompile("/ajax/libs/react-grid-layout/0.8.0"),
		glob.MustCompile("/ajax/libs/react-mdl/0.*"),
		glob.MustCompile("/ajax/libs/react-mdl/1.0.[0-1]"),
		glob.MustCompile("/ajax/libs/react-quill/02e9da4/"),
		glob.MustCompile("/ajax/libs/react-router/latest/"),
		glob.MustCompile("/ajax/libs/redux-form/6.0.3/"),
		glob.MustCompile("/ajax/libs/require-cs/latest/"),
		glob.MustCompile("/ajax/libs/require-css/dev-test/"),
		glob.MustCompile("/ajax/libs/require-css/release-test/"),
		glob.MustCompile("/ajax/libs/require-i18n/latest/"),
		glob.MustCompile("/ajax/libs/require.js/latest/"),
		glob.MustCompile("/ajax/libs/roslibjs/node*"),
		glob.MustCompile("/ajax/libs/rrssb/1.75"),
		glob.MustCompile("/ajax/libs/rrssb/1.76"),
		glob.MustCompile("/ajax/libs/ryejs/0.1.1/"),
		glob.MustCompile("/ajax/libs/ryejs/0.1.2/"),
		glob.MustCompile("/ajax/libs/single-spa/[12].*/"),
		glob.MustCompile("/ajax/libs/sbt/1.0.0.20130926-1550/"),
		glob.MustCompile("/ajax/libs/sbt/1.1.3-20150220/"),
		glob.MustCompile("/ajax/libs/sbt/IBMSBT-20140312/"),
		glob.MustCompile("/ajax/libs/sbt/sbt*-1.[0-9].[0-9]-201[0-9]*/"),
		glob.MustCompile("/ajax/libs/sbt/sb[tk]*"),
		glob.MustCompile("/ajax/libs/seajs/2.0.0b[1-3]"),
		glob.MustCompile("/ajax/libs/slim.js/1.*/"),
		glob.MustCompile("/ajax/libs/slim.js/2.[0-5].*/"),
		glob.MustCompile("/ajax/libs/slippry/v.1.2/"),
		glob.MustCompile("/ajax/libs/social-likes/path/"),
		glob.MustCompile("/ajax/libs/spoqa-han-sans/origin"),
		glob.MustCompile("/ajax/libs/starability/v.0.1.0"),
		glob.MustCompile("/ajax/libs/strophe.js/release-1.2.0"),
		glob.MustCompile("/ajax/libs/swagger-ui/v/"),
		glob.MustCompile("/ajax/libs/systemjs-plugin-json/0.1.0/"),
		glob.MustCompile("/ajax/libs/tachyons/v.4.5.0"),
		glob.MustCompile("/ajax/libs/tent-css/*/06_components"),
		glob.MustCompile("/ajax/libs/tmlib.js/rsource0.1.0/"),
		glob.MustCompile("/ajax/libs/tocas-ui/2.0.0-rc.*"),
		glob.MustCompile("/ajax/libs/trackpad-scroll-emulator/version_1.0"),
		glob.MustCompile("/ajax/libs/Turf.js/0.0.4[3-6]"),
		glob.MustCompile("/ajax/libs/tween.js/0.0.0-development/"),
		glob.MustCompile("/ajax/libs/typebase.css/0.5"),
		glob.MustCompile("/ajax/libs/typescript/*dev*"),
		glob.MustCompile("/ajax/libs/typescript/*insiders*"),
		glob.MustCompile("/ajax/libs/typescript/*/typescriptServices.js"),
		glob.MustCompile("/ajax/libs/uikit/*-dev.*"),
		glob.MustCompile("/ajax/libs/underscore.string/v.1.1.3"),
		glob.MustCompile("/ajax/libs/underscore.string/v.2.0.0"),
		glob.MustCompile("/ajax/libs/underscore.string/v.2.1.1"),
		glob.MustCompile("/ajax/libs/underscore.string/v.2.2.0rc"),
		glob.MustCompile("/ajax/libs/unibox/jquery_plugin/"),
		glob.MustCompile("/ajax/libs/URI.js/v.1*"),
		glob.MustCompile("/ajax/libs/vega-tooltip/0.3.1"),
		glob.MustCompile("/ajax/libs/vibrant.js/1.0"),
		glob.MustCompile("/ajax/libs/v-mask/1.[0-2].*"),
		glob.MustCompile("/ajax/libs/vue-lazyload/0.*"),
		glob.MustCompile("/ajax/libs/vue-lazyload/1.0.0-*"),
		glob.MustCompile("/ajax/libs/Vue.Draggable/v.19.2/"),
		glob.MustCompile("/ajax/libs/wavesurfer.js/1.0.23"),
		glob.MustCompile("/ajax/libs/wavesurfer.js/1.0.26"),
		glob.MustCompile("/ajax/libs/wavesurfer.js/marky-boy_2006-01-11/"),
		glob.MustCompile("/ajax/libs/waypoints/latest/"),
		glob.MustCompile("/ajax/libs/waypoints/latest/"),
		glob.MustCompile("/ajax/libs/weather-icons/1.1endpoint"),
		glob.MustCompile("/ajax/libs/weather-icons/1.1ghpagesendpoint"),
		glob.MustCompile("/ajax/libs/winjs/release"),
		glob.MustCompile("/ajax/libs/xterm/3.*-beta*"),
		glob.MustCompile("/ajax/libs/xterm/3.14.0"),
		glob.MustCompile("/ajax/libs/yamlcss/rm,"),
		glob.MustCompile("/ajax/libs/webvr-polyfill/1.0.0"),
		glob.MustCompile("/ajax/libs/webvr-polyfill/0.[0-9].*"),
		glob.MustCompile("/ajax/libs/firebase/*-canary.*"),
		glob.MustCompile("/ajax/libs/colresizable/1.6/"),

		// ignore snapshot tags
		glob.MustCompile("/ajax/libs/*/*snapshot*"),

		// ignore files that we don't need
		glob.MustCompile("/ajax/libs/*/*/**/Gruntfile.js"),
		glob.MustCompile("/ajax/libs/*/*/**/.gitignore"),
	}
)
