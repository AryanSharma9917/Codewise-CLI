import { ssrRenderAttrs } from "vue/server-renderer";
import { useSSRContext } from "vue";
import { _ as _export_sfc } from "./plugin-vue_export-helper.1tPrXgE0.js";
const __pageData = JSON.parse('{"title":"How Codewise works","description":"","frontmatter":{},"headers":[],"relativePath":"guide/architecture.md","filePath":"guide/architecture.md","lastUpdated":null}');
const _sfc_main = { name: "guide/architecture.md" };
function _sfc_ssrRender(_ctx, _push, _parent, _attrs, $props, $setup, $data, $options) {
  _push(`<div${ssrRenderAttrs(_attrs)}><h1 id="how-codewise-works" tabindex="-1">How Codewise works <a class="header-anchor" href="#how-codewise-works" aria-label="Permalink to &quot;How Codewise works&quot;">​</a></h1><p>Codewise is a Go application built with Cobra. Commands parse user input, domain packages implement the behavior, and external tools perform platform operations.</p><div class="language-text vp-adaptive-theme"><button title="Copy Code" class="copy"></button><span class="lang">text</span><pre class="shiki shiki-themes github-light github-dark vp-code" tabindex="0"><code><span class="line"><span>Terminal</span></span>
<span class="line"><span>  └─ codewise</span></span>
<span class="line"><span>      ├─ cmd/                 command definitions and flags</span></span>
<span class="line"><span>      └─ pkg/                 workflow logic</span></span>
<span class="line"><span>          ├─ local files     config, templates, manifests, charts</span></span>
<span class="line"><span>          └─ external CLIs   docker, kubectl, helm</span></span>
<span class="line"><span>                              └─ registry or Kubernetes API</span></span></code></pre></div><h2 id="repository-map" tabindex="-1">Repository map <a class="header-anchor" href="#repository-map" aria-label="Permalink to &quot;Repository map&quot;">​</a></h2><table tabindex="0"><thead><tr><th>Path</th><th>Responsibility</th></tr></thead><tbody><tr><td><code>main.go</code></td><td>Program entry point</td></tr><tr><td><code>cmd/</code></td><td>Cobra commands, flags, and input validation</td></tr><tr><td><code>pkg/config/</code></td><td>Global configuration</td></tr><tr><td><code>pkg/docker/</code></td><td>Dockerfile and image operations</td></tr><tr><td><code>pkg/k8s/</code></td><td>Manifest generation and Kubernetes execution</td></tr><tr><td><code>pkg/helm/</code></td><td>Helm chart generation</td></tr><tr><td><code>pkg/deploy/</code></td><td>Plans, strategies, execution, diagnostics, and rollback</td></tr><tr><td><code>pkg/env/</code></td><td>Environment profile lifecycle</td></tr><tr><td><code>pkg/encoder/</code></td><td>Data conversion utilities</td></tr><tr><td><code>pkg/generator/</code></td><td>Project and automation templates</td></tr></tbody></table><h2 id="deployment-strategy-selection" tabindex="-1">Deployment strategy selection <a class="header-anchor" href="#deployment-strategy-selection" aria-label="Permalink to &quot;Deployment strategy selection&quot;">​</a></h2><p>Codewise resolves deployment strategies in this order:</p><ol><li>GitOps when the environment has a Git repository configured.</li><li>Helm when <code>helm/chart/</code> exists.</li><li>Raw Kubernetes manifests otherwise.</li></ol><p>Use <code>codewise deploy plan --env NAME</code> to see the selected strategy before execution.</p></div>`);
}
const _sfc_setup = _sfc_main.setup;
_sfc_main.setup = (props, ctx) => {
  const ssrContext = useSSRContext();
  (ssrContext.modules || (ssrContext.modules = /* @__PURE__ */ new Set())).add("guide/architecture.md");
  return _sfc_setup ? _sfc_setup(props, ctx) : void 0;
};
const architecture = /* @__PURE__ */ _export_sfc(_sfc_main, [["ssrRender", _sfc_ssrRender]]);
export {
  __pageData,
  architecture as default
};
