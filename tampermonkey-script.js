// ==UserScript==
// @name         deploy with coding
// @namespace    http://tampermonkey.net/
// @version      0.1
// @description  shows how to use babel compiler
// @author       sqlwwx
// @require      https://cdn.staticfile.org/axios/0.15.3/axios.min.js
// @require      https://cdnjs.cloudflare.com/ajax/libs/url-search-params/0.10.0/url-search-params.js
// @require      https://cdnjs.cloudflare.com/ajax/libs/babel-standalone/6.18.2/babel.js
// @require      https://cdnjs.cloudflare.com/ajax/libs/babel-polyfill/6.16.0/polyfill.js
// @match        https://coding.net/t/*/p/*/git
// ==/UserScript==

/* jshint ignore:start */
var inline_src = (<><![CDATA[
    /* jshint ignore:end */
    /* jshint esnext: false */
    /* jshint esversion: 6 */
    $(() => {
      const secret = 'mysecret';
      setTimeout(()=> {
        $('.nav-breadcrumb:first').append('<button id="deploy-test">部署测试环境</button>');
        $('.nav-breadcrumb:first').append('<button id="deploy-prod">部署生产环境</button>');
        const paths = window.location.pathname.split('\/');
        const nameSpace = paths[2];
        const project = paths[4];
        const host = 'https://deployer.lab.wuweixing.com';
        function deploy (env) {
            const id = [nameSpace, project].join('_');
            var params = new URLSearchParams();
            params.append('secret', secret);
            axios.post(
                host + '/' + id,
                 params,
                {withCredentials:true}
            ).then((res) => {
                if (res.data === 'success') {
                    window.alert("部署成功");
                } else {
                    window.alert("部署失败");
                }
            }).catch(function (error) {
                console.log(error);
                window.alert("部署失败", error);
            });
        }
        $('#deploy-test').click(() => { deploy('test');});
        $('#deploy-prod').click(() => { deploy('prod');});
      }, 3000);
    });
    /* jshint ignore:start */
]]></>).toString();
                  var c = Babel.transform(inline_src, { presets: [ "es2015", "es2016" ] });
eval(c.code);
/* jshint ignore:end */
