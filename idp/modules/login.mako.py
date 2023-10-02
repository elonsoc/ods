# -*- coding:utf-8 -*-
from mako import runtime, filters, cache
UNDEFINED = runtime.UNDEFINED
STOP_RENDERING = runtime.STOP_RENDERING
__M_dict_builtin = dict
__M_locals_builtin = locals
_magic_number = 10
_modified_time = 1694530231.8080142
_enable_loop = True
_template_filename = 'htdocs/login.mako'
_template_uri = 'login.mako'
_source_encoding = 'utf-8'
_exports = []


def _mako_get_namespace(context, name):
    try:
        return context.namespaces[(__name__, name)]
    except KeyError:
        _mako_generate_namespaces(context)
        return context.namespaces[(__name__, name)]
def _mako_generate_namespaces(context):
    pass
def _mako_inherit(template, context):
    _mako_generate_namespaces(context)
    return runtime._inherit_from(context, 'root.mako', _template_uri)
def render_body(context,**pageargs):
    __M_caller = context.caller_stack._push_frame()
    try:
        __M_locals = __M_dict_builtin(pageargs=pageargs)
        action = context.get('action', UNDEFINED)
        key = context.get('key', UNDEFINED)
        redirect_uri = context.get('redirect_uri', UNDEFINED)
        authn_reference = context.get('authn_reference', UNDEFINED)
        __M_writer = context.writer()
        __M_writer('\n\n<h1>Please log in</h1>\n<p class="description">\n    To register it\'s quite simple: enter a login and a password\n</p>\n\n<form action="')
        __M_writer(str(action))
        __M_writer('" method="post">\n    <input type="hidden" name="key" value="')
        __M_writer(str(key))
        __M_writer('"/>\n    <input type="hidden" name="authn_reference" value="')
        __M_writer(str(authn_reference))
        __M_writer('"/>\n    <input type="hidden" name="redirect_uri" value="')
        __M_writer(str(redirect_uri))
        __M_writer('"/>\n\n    <div class="label">\n        <label for="login">Username</label>\n    </div>\n    <div>\n        <input type="text" name="login" value="emusk" autofocus><br/>\n    </div>\n\n    <div class="label">\n        <label for="password">Password</label>\n    </div>\n    <div>\n        <input type="password" name="password"\n               value="elonmuskisbest"/>\n    </div>\n\n    <input class="submit" type="submit" name="form.submitted" value="Log In"/>\n</form>\n')
        return ''
    finally:
        context.caller_stack._pop_frame()


"""
__M_BEGIN_METADATA
{"filename": "htdocs/login.mako", "uri": "login.mako", "source_encoding": "utf-8", "line_map": {"27": 0, "36": 1, "37": 8, "38": 8, "39": 9, "40": 9, "41": 10, "42": 10, "43": 11, "44": 11, "50": 44}}
__M_END_METADATA
"""
