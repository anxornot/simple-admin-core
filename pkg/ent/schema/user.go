package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/suyuan32/simple-admin-core/pkg/ent/schema/mixins"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").Unique().Comment("User's login name | 登录名"),
		field.String("password").Comment("Password | 密码"),
		field.String("nickname").Unique().Comment("Nickname | 昵称"),
		field.String("description").Optional().Comment("The description of user | 用户的描述信息"),
		field.String("home_path").Default("/dashboard").Comment("The home page that the user enters after logging in | 用户登陆后进入的首页"),
		field.Uint64("role_id").Optional().Default(2).Comment("Role id | 角色ID"),
		field.String("mobile").Optional().Comment("Mobile number | 手机号"),
		field.String("email").Optional().Comment("Email | 邮箱号"),
		field.String("avatar").
			SchemaType(map[string]string{dialect.MySQL: "varchar(512)"}).
			Optional().
			Default("").
			Comment("Avatar | 头像路径"),
		field.Uint64("department_id").Optional().Default(1).Comment("Department ID | 部门ID"),
		field.Uint64("position_id").Optional().Default(1).Comment("Position ID | 职位ID"),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.UUIDMixin{},
		mixins.StatusMixin{},
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("department", Department.Type).Unique().Field("department_id"),
		edge.To("position", Position.Type).Unique().Field("position_id"),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("username", "email").
			Unique(),
	}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_users"},
	}
}
