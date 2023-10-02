# -*- coding:utf-8 -*-
from mako import runtime, filters, cache
UNDEFINED = runtime.UNDEFINED
STOP_RENDERING = runtime.STOP_RENDERING
__M_dict_builtin = dict
__M_locals_builtin = locals
_magic_number = 10
_modified_time = 1693674406.6983511
_enable_loop = True
_template_filename = 'templates/root.mako'
_template_uri = 'root.mako'
_source_encoding = 'utf-8'
_exports = ['css_link', 'css', 'pre', 'post']


def render_body(context,**pageargs):
    __M_caller = context.caller_stack._push_frame()
    try:
        __M_locals = __M_dict_builtin(pageargs=pageargs)
        next = context.get('next', UNDEFINED)
        set = context.get('set', UNDEFINED)
        def pre():
            return render_pre(context._locals(__M_locals))
        def post():
            return render_post(context._locals(__M_locals))
        self = context.get('self', UNDEFINED)
        __M_writer = context.writer()
        self.seen_css = set() 
        
        __M_writer('\n')
        __M_writer('\n')
        __M_writer('\n')
        __M_writer('\n')
        __M_writer('\n')
        __M_writer('<html>\n<head><title>IDP test login</title>\n    ')
        __M_writer(str(self.css()))
        __M_writer('\n    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />\n</head>\n<body>\n    ')
        __M_writer(str(pre()))
        __M_writer('\n')
        __M_writer(str(next.body()))
        __M_writer('\n')
        __M_writer(str(post()))
        __M_writer('\n</body>\n</html>\n')
        return ''
    finally:
        context.caller_stack._pop_frame()


def render_css_link(context,path,media=''):
    __M_caller = context.caller_stack._push_frame()
    try:
        context._push_buffer()
        self = context.get('self', UNDEFINED)
        __M_writer = context.writer()
        __M_writer('\n')
        if path not in self.seen_css:
            __M_writer('        <link rel="stylesheet" type="text/css" href="')
            __M_writer(filters.html_escape(str(path)))
            __M_writer('" media="')
            __M_writer(str(media))
            __M_writer('">\n')
        __M_writer('    ')
        self.seen_css.add(path) 
        
        __M_writer('\n')
    finally:
        __M_buf, __M_writer = context._pop_buffer_and_writer()
        context.caller_stack._pop_frame()
    __M_writer(filters.trim(__M_buf.getvalue()))
    return ''


def render_css(context):
    __M_caller = context.caller_stack._push_frame()
    try:
        context._push_buffer()
        def css_link(path,media=''):
            return render_css_link(context,path,media)
        __M_writer = context.writer()
        __M_writer('\n    ')
        __M_writer(str(css_link('/static/css/main.css', 'screen')))
        __M_writer('\n')
    finally:
        __M_buf, __M_writer = context._pop_buffer_and_writer()
        context.caller_stack._pop_frame()
    __M_writer(filters.trim(__M_buf.getvalue()))
    return ''


def render_pre(context):
    __M_caller = context.caller_stack._push_frame()
    try:
        context._push_buffer()
        __M_writer = context.writer()
        __M_writer('\n    <div class="header">\n        <h1><a href="/">Login</a></h1>\n    </div>\n')
    finally:
        __M_buf, __M_writer = context._pop_buffer_and_writer()
        context.caller_stack._pop_frame()
    __M_writer(filters.trim(__M_buf.getvalue()))
    return ''


def render_post(context):
    __M_caller = context.caller_stack._push_frame()
    try:
        context._push_buffer()
        __M_writer = context.writer()
        __M_writer('\n    <div>\n        <div class="footer">\n            <p>&#169; Copyright 2014 Ume&#229; Universitet &nbsp;</p>\n        </div>\n    </div>\n')
    finally:
        __M_buf, __M_writer = context._pop_buffer_and_writer()
        context.caller_stack._pop_frame()
    __M_writer(filters.trim(__M_buf.getvalue()))
    return ''


"""
__M_BEGIN_METADATA
{"filename": "templates/root.mako", "uri": "root.mako", "source_encoding": "utf-8", "line_map": {"16": 0, "28": 1, "29": 2, "30": 1, "31": 7, "32": 10, "33": 15, "34": 22, "35": 25, "36": 27, "37": 27, "38": 31, "39": 31, "40": 34, "41": 34, "42": 35, "43": 35, "49": 2, "55": 2, "56": 3, "57": 4, "58": 4, "59": 4, "60": 4, "61": 4, "62": 6, "63": 6, "64": 7, "65": 6, "73": 8, "80": 8, "81": 9, "82": 9, "90": 11, "95": 11, "103": 16, "108": 16, "116": 108}}
__M_END_METADATA
"""
