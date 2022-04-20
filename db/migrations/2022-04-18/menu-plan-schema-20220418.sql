CREATE TABLE public."comments"
(
    comment_id text        NOT NULL,
    "comment"  text        NOT NULL,
    user_id    text        NOT NULL,
    created_at timestamptz NOT NULL,
    updated_at timestamptz NOT NULL,
    CONSTRAINT comments_pkey PRIMARY KEY (comment_id)
);

CREATE TABLE public.menus
(
    menu_id            text        NOT NULL,
    menu_name          text        NOT NULL,
    menu_description   text        NOT NULL,
    total_cocking_time text        NOT NULL,
    week               int8        NOT NULL,
    created_at         timestamptz NOT NULL,
    updated_at         timestamptz NOT NULL,
    CONSTRAINT menus_pkey PRIMARY KEY (menu_id)
);


CREATE TABLE public.users
(
    user_id    text        NOT NULL,
    user_name  text        NOT NULL,
    email      text        NOT NULL,
    user_type  text        NOT NULL,
    "password" text        NOT NULL,
    created_at timestamptz NOT NULL,
    updated_at timestamptz NOT NULL,
    CONSTRAINT users_pkey PRIMARY KEY (user_id)
);


CREATE TABLE public.recipes
(
    recipe_id             text        NOT NULL,
    recipe_name           text        NOT NULL,
    recipe_description    text NULL,
    recipe_instructions   text NULL,
    menu_id               text NULL,
    recipe_classification text NULL,
    created_at            timestamptz NOT NULL,
    updated_at            timestamptz NOT NULL,
    CONSTRAINT recipes_pkey PRIMARY KEY (recipe_id),
    CONSTRAINT fk_menus_recipe FOREIGN KEY (menu_id) REFERENCES public.menus (menu_id) ON DELETE CASCADE
);


CREATE TABLE public.ingredients
(
    ingredient_id          text        NOT NULL,
    ingredient_name        text        NOT NULL,
    ingredient_description text        NOT NULL,
    unit                   int8        NOT NULL,
    recipe_id              text NULL,
    created_at             timestamptz NOT NULL,
    updated_at             timestamptz NOT NULL,
    CONSTRAINT ingredients_pkey PRIMARY KEY (ingredient_id),
    CONSTRAINT fk_recipes_ingredient FOREIGN KEY (recipe_id) REFERENCES public.recipes (recipe_id) ON DELETE CASCADE
);