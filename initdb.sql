create table accounts(
    id              bigserial primary key,
    document_number varchar(11) not null unique,
    created_at      timestamp default now(),
    created_from      varchar(15) not null
);

CREATE UNIQUE INDEX document_number_idx ON accounts (document_number);

create table operation_types(
    id           bigint primary key,
    description  varchar(30) not null
);

create table transactions(
    id                  bigserial primary key,
    account_id          bigint not null,
    operation_type_id   bigint not null,
    amount              numeric(12, 2) not null,
    created_at          timestamp default now(),
    created_from          varchar(15) not null,
    foreign key (account_id) references accounts(id) 
        on delete cascade,
    foreign key (operation_type_id) references operation_types(id)
);

insert into operation_types(id, description) 
    values 
     (1, 'COMPRA A VISTA')
    ,(2, 'COMPRA PARCELADA')
    ,(3, 'SAQUE')
    ,(4, 'PAGAMENTO');

