// This code is transcompiled by LuneScript.
package lns
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
import LnsTypes "github.com/ifritJP/LuneScript/src/lune/base"
var init_EbnfDef bool
var EbnfDef__mod__ string
var EbnfDef_ebnf string
// 243: decl @lns.@EbnfDef.getEbnfSrc
func EbnfDef_getEbnfSrc(_env *LnsEnv) LnsAny {
    return &LnsTypes.Types_TokenizerSrc__LnsCode{EbnfDef_ebnf, "", nil}
}

func Lns_EbnfDef_init(_env *LnsEnv) {
    if init_EbnfDef { return }
    init_EbnfDef = true
    EbnfDef__mod__ = "@lns.@EbnfDef"
    Lns_InitMod()
    LnsTypes.Lns_Types_init(_env)
    EbnfDef_ebnf = "  <comment> ::= \"//\" <anytoken_br> | \"/*\" <anytoken> \"*/\"\n\n     <code> ::= [<shebang>] {<code_header>} {<stmt>} <eof>\n\n  <code_header> ::= <subfile> | <lune_control> | <import>\n\n  <shebang> ::= \"#!\" <token> <anytoken_br>\n\n  <import> ::= \"import\" <lazy_direct> [ \"go\" \"/\" ] <sym> { <import_dlmt> <sym> } [ \"as\" <sym> ] \";\"\n  <import_dlmt> ::= \".\" | \":\"\n\n\n  <subfile> ::= <subfile_owner> | <subfile_sub>\n  <subfile_owner> ::= \"subfile\" \"owner\" <module_path> \";\"\n  <subfile_sub> ::= { \"subfile\" \"use\" <module_path> \";\" }\n  <module_path> ::= <sym> { \".\" <sym> }\n\n  <stmt> ::= <decl_access> <declaration> | <test> | <block> | <super> |\n\t   <asynclock> | <if> | <if_unwrap> | <when> | <switch> | <match> | <while> |\n\t     <repeat> | <for> | <apply> | <foreach> | <forsort> | <return> |\n\t     <break> | <unwrap_stmt> | <scope> | <lune_control> |\n             <test_block> |\n\t     <none> | <stmt_exp>\n\n  <test_block> ::= \"__test\" [<sym> \"(\" <sym> \")\" ] \"{\" {<import>} { <stmt> } \"}\"\n\n  <declaration> ::= <decl_var> | <decl_fn> | \n\t\t    <decl_class> | <decl_interface> | <decl_module> |\n\t\t    <decl_proto> | <decl_macro> | <alge> | <form> | <alias> | <decl_enum>\n\n\n  <decl_access> ::= [<access>] [\"static\"]\n  <decl_macro> ::= \"macro\" <sym> <decl_arg_paren>\n\t\t   [ \":\" <reftype_list> ] \"{\" [ <block> ] <stat> \"}\"\n  <fn_attrib> ::= [<decl_async>] [\"mut\"]\n  <form> ::= \"form\" <sym> [<generics>] <decl_arg_paren>\n\t\t<fn_attrib> [\":\" <reftype_list>] \";\"\n  <decl_fn> ::= [\"override\"] \"fn\" [ <sym> \".\" ] <sym> [<generics>] <decl_arg_paren>\n\t\t<fn_attrib> [\":\" <reftype_list>] <decl_fn_body>\n  <decl_fn_body> ::= <block> | \";\"\n\n  <decl_class_attrib> ::= [ \"abstract\" ] [ \"final\" ]\n\n  <decl_class> ::= <decl_class_attrib> \"class\" <sym> [<generics>] \n\t\t   [ \"extend\" [<type>] [\"(\"<type_list> \")\" ] ] \"{\" { <class_field> } \"}\"\n  <decl_proto> ::= \"proto\" <decl_class_attrib> <decl_proto_class_if> <sym> [<generics>] \n\t\t   [ \"extend\" [<type>] [\"(\"<type_list> \")\" ] ] \";\"\n  <decl_proto_class_if> ::= \"class\" | \"interface\"\n  <decl_interface> ::= \"interface\" <sym> [<generics>] \n\t\t       [ \"extend\" <type> ] \"{\" { <decl_access> <if_field> } \"}\"\n  <decl_module> ::= \"module\" <lazy_direct> <sym> \"require\" <literal_str>\n\t\t       \"{\" { <decl_access> <module_field> } \"}\"\n  <lazy_direct> ::= [ \".\" \"l\" ] | [ \".\" \"d\" ]\n  <alge> ::= \"alge\" <sym> [<generics>] \"{\" { <alge_val> [\",\"] } \"}\"\n  <alias> ::= \"alias\" <sym> \"=\" <reftype> \";\"\n  <decl_enum> ::= \"enum\" <sym> \"{\" {<enum_val> [ \",\" ] } \"}\"\n  <decl_member> ::= \"let\" <decl_sym> [ \"{\" <accessor> [\",\" <accessor>] \"}\" ] \";\"\n  <decl_method> ::= [ \"abstract\" ] [\"override\"] \"fn\" <sym> [<generics>]\n\t\t    <decl_arg_paren> <fn_attrib> [\":\" \n\t\t    <reftype_list>] <decl_method_body>\n  <decl_var> ::= \"let\" <decl_sym_list> [ \"=\" <exp_list> ] \";\" |\n\t       \"let\" \"!\" <decl_sym_list> \"=\" <exp_list> <block> [ \"then\" <block> ] [ \"else\" <block> ] \";\"\n\n  <decl_async> ::= \"__async\" | \"__noasync\" | \"__trans\"\n\n  <test> ::= \"__test\" <sym> \"(\" <sym> \")\" <block>\n\n  <block> ::= \"{\" { <stmt> } \"}\"\n  <super> ::= \"super\" <arg_list> \";\"\n\n  <arg_paren_begin> ::= \"(\" | \"$(\"\n  <arg_list> ::=<arg_paren_begin> [ <arg> {\",\" <arg>} ] [\"##\"] \")\"\n  <arg> ::= <exp>\n\n  <asynclock> ::= \"__asyncLock\" <block> |\n\t\t  \"__luago\" <block> | \n\t\t  \"__luaLock\" <block> | \n\t\t  \"__luaDepend\" <block>\n  <if> ::= <if_base> { <if_elseif> } [ <if_else> ]\n  <if_base> ::= \"if\" <exp> <block>\n  <if_elseif> ::= \"elseif\" <exp> <block>\n  <if_else> ::= \"else\" <block>\n  <if_unwrap> ::= <if_unwrap_base> [ <if_else> ]\n  <if_unwrap_base> ::= \"if\" \"!\" [ \"let\" <decl_sym_list> \"=\" ] <exp_list> <block>\n  <sym_list> ::= <sym> { \",\" <sym> }\n  <when> ::= \"when\" \"!\" <exp_list> <block> [ \"else\" <block> ]\n  <switch> ::= <switch_keyword> <exp> \"{\" { <switch_case> }\n\t       [ <default> ] \"}\"\n  <switch_case> ::= \"case\" <exp_list> <block>\n  <switch_keyword> ::= \"switch\" | \"_switch\"\n  <default> ::= \"default\" <block>\n\n\n  <exp_list> ::= <exp> { \",\" <exp> }\n  <match> ::= <match_keyword> <exp> \"{\" { <match_case> } [ <default> ] \"}\"\n  <match_case> ::= \"case\" <match_case_val> <block>\n  <match_case_val> ::= [ <type> ] \".\" <sym> [\"(\" <sym_list> \")\"]\n  <match_keyword> ::= \"match\" | \"_match\"\n  <sym_list> ::= <sym> { \",\" <exp> }\t      \n  <while> ::= \"while\" <exp> <block>\n  <repeat> ::= \"repeat\" <block> <exp> \";\"\n  <for> ::= \"for\" <sym> \"=\" <exp> \",\" <exp> [ \",\" <exp> ] <block>\n  <apply> ::= \"apply\" <sym_list> \"of\" <exp> <block>\n  <foreach> ::= \"foreach\" <sym_list> \"in\" <exp> <block>\n  <forsort> ::= \"forsort\" <sym_list> \"in\" <exp> <block>\n  <return> ::= \"return\" [<exp_list>] \";\"\n  <break> ::= \"break\" \";\"\n  <unwrap_stmt> ::= \"unwrap\" \"!\" <sym_list> \"=\" <exp_list> <block> \";\"\n  <scope> ::= \"__scope\" \"root\" \"(\" <sym_list> \")\" <block>\n  <lune_control> ::= \"_lune_control\" <sym> <anytoken> \";\"\n  <none> ::= \";\"\n  <stmt_exp> ::= <exp_list> [ \"=\" <exp_list> ] \";\"\n\n  <anytoken> ::= { <token> }\n\n\n\n  <decl_sym_list> ::= <decl_sym> { \",\" <decl_sym> }\n  <decl_sym> ::= [<decl_sym_mut>] <sym> [\":\" <membertype>]\n  <decl_sym_mut> ::= \"mut\" | \"allmut\"\n\n  <membertype> ::= [ \"allmut\" ] <reftype>\n\n\n  <reftype_list> ::= <reftype> { \",\" <reftype> }\n\n  <decl_arg_paren> ::= \"(\" [ <decl_arg_list> ] \")\"\n  <decl_arg_list> ::= <decl_arg> { \",\" <decl_arg> }\n\n  <decl_arg> ::= [\"mut\"] <sym> \":\" <reftype> | <decl_arg_ddd>\n  <decl_arg_ddd> ::= \"...\" [ \"<\" <reftype> \">\"]\n\n  <type_list> ::= <type> { \",\" <type> }\n  <type> ::= <sym> { \".\" <sym> } [<generics>]\n\n  <class_field> ::= <decl_access> <decl_field> | <init_block> | <advertise> |\n\t\t    <lune_control> | <exp_macro> | <none>\n\n  <decl_field> ::= <decl_member> | <decl_method> | <decl_enum> \n\n  <init_block> ::= \"__init\" <block>\n\n  <exp_macro> ::= <sym> \"(\" <exp_list> \")\"\n\n\n\n  <if_field> ::= <decl_access> <decl_method>\n\n  <module_field> ::= <decl_member> | <decl_method> \n\n  <generics> ::= \"<\" <generics_decl> {\",\" <generics_decl>} \">\"\n\n  <generics_decl_one> ::= <reftype> [ \":\" [\"(\"] <reftype> [\")\"] ]\n  <generics_decl> ::= <sym> \"=\" <generics_decl_one> | <generics_decl_one>\n\n  <advertise> ::= \"advertise\" <sym> \";\"\n\n\n  <alge_val> ::= <sym> [ \"(\" { <alge_val_param> [ \",\" ]  }\")\" ]\n\n  <alge_val_param> ::= <sym> \":\" <reftype> | <reftype>\n\n  <decl_tuple> ::= \"(\" <decl_tuple_item> {\",\" <decl_tuple_item> } \")\"\n  <decl_tuple_item> ::= <sym> \":\" <reftype> | <reftype>\n\n  <reftype> ::= [\"&\"] <reftype_raw> [<reftype_suffix>] [\"!\"]\n  <reftype_raw> ::= <decl_tuple> | <type> [<set_generics>]\n  <reftype_suffix> ::= \"[\" \"]\" | \"[@\" \"]\"\n\n\n  <set_generics> ::= \"<\" <reftype> { \",\" <reftype> } \">\"\n\n  <access> ::= \"global\" | \"pub\" | \"pro\" | \"pri\" | \"local\"\n\n\n  <enum_val> ::= <sym> [ \"=\" <exp> ]\n\n\n  <accessor> ::= <accessor_kind> [\"&\"] [ \":\" <reftype> ]\n  <accessor_kind> ::= \"non\" | \"pub\" | \"pri\" | \"pro\" | \"local\"\n\n  <decl_method_body> ::= <block> | \";\"\n\n\n\n  <exp> ::= <exp_wrap>\n  <exp_wrap> ::= [<op1>] <exp_single> { <op2> <exp_wrap> }\n  <exp_single> ::= <exp_one> { <exp_suffix> }\n  <exp_one> ::= \"...\" | <omit_enum> | <const_list> | <const_array> | <const_set> | \n\t\t<const_tuple> | <const_map> |\n\t\t<paren> | <new> | <literal_int> | <literal_real > |\n\t\t<literal_char> | <const_str> | \n\t\t<unwrap> | <anonymous_func> | \n\t\t<literal_bool> | \"nil\" | \"null\" | \"##\" | \n                <stat_block> | <block> | <reftype> | <ref_sym>\n  <stat_block> ::= \"`{\" <stat> \"}\"\n  <const_str> ::= <literal_str> [ <arg_list>]\n  <exp_suffix> ::= [<ref_sym_dlmt> <exp_one>] [ <ref_index> ] [ <call> ] \n\t\t   [ \"!\" ] [\"...\" [ \"**\"] ] [<cast>]\n  <ref_index_begin> ::= \"[\" | \"$[\"\n  <ref_index> ::= <ref_index_begin> <exp> \"]\"\n  <call> ::= [<set_generics>] <arg_list> [\"**\"]\n\n  <new> ::= \"new\" <type> [<set_generics>] <arg_list>\n\n  <omit_enum> ::= \".\" <sym>\n  <paren> ::= \"(\" <exp> \")\"\n  <ref_sym> ::= <sym> { <ref_sym_dlmt> <sym> }\n  <ref_sym_dlmt> ::= \".\" | \".$\" | \"$.\" | \"$.$\"\n\n  <cast> ::= <cast_op> <reftype>\n  <cast_op> ::= \"@@\" | \"@@=\" | \"@@@\"\n\n  <anonymous_func> ::= \"fn\" [<generics>] <decl_arg_paren>\n\t\t       <fn_attrib> [\":\" <reftype_list>] <block>\n\n\n  <unwrap> ::= \"unwrap\" <exp> [default <exp>]\n\n  <const_list> ::= \"[\" [<exp_list>] \"]\"\n  <const_map> ::= \"{\" {<const_map_entry>} \"}\"\n  <const_map_entry> ::= <exp> \":\" <exp> [ \",\" ]\n  <const_set> ::= \"(@\" [<exp_list>] \")\"\n  <const_array> ::= \"[@\" [<exp_list>] \"]\"\n  <const_tuple> ::= \"(=\" [<exp_list>] \")\"\n\n  <literal_bool> ::= true | false\n\n  <op1> ::= \"not\" | \"#\" | \"~\" | \",,\" | \",,,\" | \",,,,\"\n\n  <op2> ::= \"+\" | \"-\" | \"*\" | \"/\" | \"^\" | \"%\" | \"&\" | \"~\" |\n\t    \"|\" | \"|>>\" | \"|<<\" | \"..\" |\n\t    \"<\" | \"<=\" | \">\" | \">=\" | \"==\" | \"~=\" | \"and\" | \"or\"\n\n"
}
func init() {
    init_EbnfDef = false
}