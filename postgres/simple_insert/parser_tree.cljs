:commandType

3

:querySource

0

:canSetTag

true

:utilityStmt

<>

:resultRelation

1

:hasAggs

false

:hasWindowFuncs

false

:hasTargetSRFs

false

:hasSubLinks

false

:hasDistinctOn

false

:hasRecursive

false

:hasModifyingCTE

false

:hasForUpdate

false

:hasRowSecurity

false

:cteList

<>

:rtable

({RTE
  :alias
  <>
  :eref
  {ALIAS
   :aliasname
   users
   :colnames
   ("id" "age" "name" "surname")}
  :rtekind
  0
  :relid
  16392
  :relkind
  r
  :rellockmode
  3
  :tablesample
  <>
  :lateral
  false
  :inh
  false
  :inFromCl
  false
  :requiredPerms
  1
  :checkAsUser
  0
  :selectedCols
  (b)
  :insertedCols
  (b 8 9 10)
  :updatedCols
  (b)
  :extraUpdatedCols
  (b)
  :securityQuals
  <>}
 {RTE
  :alias
  <>
  :eref
  {ALIAS
   :aliasname
   *VALUES*
   :colnames
   ("column1" "column2" "column3")}
  :rtekind
  5
  :values_lists
  (({FUNCEXPR
     :funcid
     481
     :funcresulttype
     20
     :funcretset
     false
     :funcvariadic
     false
     :funcformat
     2
     :funccollid
     0
     :inputcollid
     0
     :args
     ({CONST
       :consttype
       23
       :consttypmod
       -1
       :constcollid
       0
       :constlen
       4
       :constbyval
       true
       :constisnull
       false
       :location
       42
       :constvalue
       4 [1 0 0 0 0 0 0 0]})
     :location
     -1}
    {FUNCEXPR
     :funcid
     481
     :funcresulttype
     20
     :funcretset
     false
     :funcvariadic
     false
     :funcformat
     2
     :funccollid
     0
     :inputcollid
     0
     :args
     ({CONST
       :consttype
       23
       :consttypmod
       -1
       :constcollid
       0
       :constlen
       4
       :constbyval
       true
       :constisnull
       false
       :location
       45
       :constvalue
       4 [20 0 0 0 0 0 0 0]})
     :location
     -1}
    {FUNCEXPR
     :funcid
     669
     :funcresulttype
     1043
     :funcretset
     false
     :funcvariadic
     false
     :funcformat
     2
     :funccollid
     100
     :inputcollid
     100
     :args
     ({CONST
       :consttype
       1043
       :consttypmod
       -1
       :constcollid
       100
       :constlen
       -1
       :constbyval
       false
       :constisnull
       false
       :location
       49
       :constvalue
       7 [28 0 0 0 66 111 98]}
      {CONST
       :consttype
       23
       :consttypmod
       -1
       :constcollid
       0
       :constlen
       4
       :constbyval
       true
       :constisnull
       false
       :location
       -1
       :constvalue
       4 [3 1 0 0 0 0 0 0]}
      {CONST
       :consttype
       16
       :consttypmod
       -1
       :constcollid
       0
       :constlen
       1
       :constbyval
       true
       :constisnull
       false
       :location
       -1
       :constvalue
       1 [0 0 0 0 0 0 0 0]})
     :location
     -1})
   ({FUNCEXPR
     :funcid
     481
     :funcresulttype
     20
     :funcretset
     false
     :funcvariadic
     false
     :funcformat
     2
     :funccollid
     0
     :inputcollid
     0
     :args
     ({CONST
       :consttype
       23
       :consttypmod
       -1
       :constcollid
       0
       :constlen
       4
       :constbyval
       true
       :constisnull
       false
       :location
       58
       :constvalue
       4 [2 0 0 0 0 0 0 0]})
     :location
     -1}
    {FUNCEXPR
     :funcid
     481
     :funcresulttype
     20
     :funcretset
     false
     :funcvariadic
     false
     :funcformat
     2
     :funccollid
     0
     :inputcollid
     0
     :args
     ({CONST
       :consttype
       23
       :consttypmod
       -1
       :constcollid
       0
       :constlen
       4
       :constbyval
       true
       :constisnull
       false
       :location
       61
       :constvalue
       4 [22 0 0 0 0 0 0 0]})
     :location
     -1}
    {FUNCEXPR
     :funcid
     669
     :funcresulttype
     1043
     :funcretset
     false
     :funcvariadic
     false
     :funcformat
     2
     :funccollid
     100
     :inputcollid
     100
     :args
     ({CONST
       :consttype
       1043
       :consttypmod
       -1
       :constcollid
       100
       :constlen
       -1
       :constbyval
       false
       :constisnull
       false
       :location
       65
       :constvalue
       9 [36 0 0 0 65 108 105 99 101]}
      {CONST
       :consttype
       23
       :consttypmod
       -1
       :constcollid
       0
       :constlen
       4
       :constbyval
       true
       :constisnull
       false
       :location
       -1
       :constvalue
       4 [3 1 0 0 0 0 0 0]}
      {CONST
       :consttype
       16
       :consttypmod
       -1
       :constcollid
       0
       :constlen
       1
       :constbyval
       true
       :constisnull
       false
       :location
       -1
       :constvalue
       1 [0 0 0 0 0 0 0 0]})
     :location
     -1}))
  :coltypes
  (o 20 20 1043)
  :coltypmods
  (i -1 -1 259)
  :colcollations
  (o 0 0 0)
  :lateral
  false
  :inh
  false
  :inFromCl
  true
  :requiredPerms
  0
  :checkAsUser
  0
  :selectedCols
  (b)
  :insertedCols
  (b)
  :updatedCols
  (b)
  :extraUpdatedCols
  (b)
  :securityQuals
  <>})

:jointree

{FROMEXPR
 :fromlist
 ({RANGETBLREF
   :rtindex 2})
 :quals
 <>}

:targetList

({TARGETENTRY
  :expr
  {VAR
   :varno
   2
   :varattno
   1
   :vartype
   20
   :vartypmod
   -1
   :varcollid
   0
   :varlevelsup
   0
   :varnosyn
   2
   :varattnosyn
   1
   :location
   -1}
  :resno
  1
  :resname
  id
  :ressortgroupref
  0
  :resorigtbl
  0
  :resorigcol
  0
  :resjunk
  false}
 {TARGETENTRY
  :expr
  {VAR
   :varno
   2
   :varattno
   2
   :vartype
   20
   :vartypmod
   -1
   :varcollid
   0
   :varlevelsup
   0
   :varnosyn
   2
   :varattnosyn
   2
   :location
   -1}
  :resno
  2
  :resname
  age
  :ressortgroupref
  0
  :resorigtbl
  0
  :resorigcol
  0
  :resjunk
  false}
 {TARGETENTRY
  :expr
  {VAR
   :varno
   2
   :varattno
   3
   :vartype
   1043
   :vartypmod
   259
   :varcollid
   0
   :varlevelsup
   0
   :varnosyn
   2
   :varattnosyn
   3
   :location
   -1}
  :resno
  3
  :resname
  name
  :ressortgroupref
  0
  :resorigtbl
  0
  :resorigcol
  0
  :resjunk
  false})

:override

0

:onConflict

<>

:returningList

<>

:groupClause

<>

:groupingSets

<>

:havingQual

<>

:windowClause

<>

:distinctClause

<>

:sortClause

<>

:limitOffset

<>

:limitCount

<>

:limitOption

0

:rowMarks

<>

:setOperations

<>

:constraintDeps

<>

:withCheckOptions

<>

:stmt_location

0

:stmt_len

0

}
